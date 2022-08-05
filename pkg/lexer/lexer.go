package lexer

import (
	"vila/pkg/errorhandler"
	"vila/pkg/token"
)

type Lexer struct {
	Errors *errorhandler.ErrorList

	input           []rune
	position        int
	readPosition    int
	ch              rune
	isPreviousIdent bool
	tempNextToken   token.Token
	line            int
	column          int
}

func New(input string, errors *errorhandler.ErrorList) *Lexer {
	l := &Lexer{
		input:           []rune(input),
		Errors:          errors,
		readPosition:    0,
		line:            1,
		column:          0,
		isPreviousIdent: false,
		tempNextToken:   token.Token{},
	}

	l.readChar()
	return l
}

func (l *Lexer) AdvanceToken() token.Token {
	if l.tempNextToken.Type != "" {
		nextToken := l.tempNextToken
		l.tempNextToken = token.Token{}
		return nextToken
	}

	l.skipWhiteSpace()

	var tok token.Token

	if tok = l.lookupToken(); tok.Type != token.Illegal {
		if tok.Type == token.SlashSlash {
			l.skipComment()
			return l.AdvanceToken()
		}
		l.readChar()
		return tok
	}

	switch l.ch {
	case '"':
		tok = l.newToken(token.String, l.consumeString())
	case '\n':
		l.line += 1
		l.column = 0
		tok = l.newSingleToken(token.Endline)
		tok.Literal = l.consumeSpace()
		tok.Column = 1
	case 0:
		tok = l.newToken(token.EOF, []rune{})
	default:
		if isDigit(l.ch) {
			tokenType := token.TokenType(token.Int)
			literal := l.consumeNumber()

			if l.ch == '.' && isDigit(l.peekChar()) {
				tokenType = token.Real
				literal = append(literal, l.ch)
				l.readChar()
				literal = append(literal, l.consumeNumber()...)
			}
			tok = l.newToken(tokenType, literal)
			return tok

		} else if isLetter(l.ch) {
			tokenLiteral := l.consumeIdent()
			tokenType := token.LookupKeyword(tokenLiteral)
			tok = l.newToken(tokenType, tokenLiteral)

			// Handle spacing identifier
			if tokenType == token.Ident && !l.isPreviousIdent {

				l.isPreviousIdent = true
				nextToken := l.AdvanceToken()

				for nextToken.Type == token.Ident {
					appendToken(&tok, nextToken)
					nextToken = l.AdvanceToken()
				}

				// handle "còn nếu" token
				if nextToken.Type == token.If && (string(tok.Literal) == "còn" || string(tok.Literal) == "con") {
					appendToken(&tok, nextToken)
					l.isPreviousIdent = false
					return tok
				}

				l.isPreviousIdent = false
				l.tempNextToken = nextToken
			}

			return tok

		} else {
			l.Errors.AddLexerError("Ký tự `"+string(l.ch)+"` không hợp lệ", token.Token{
				Type:    token.Illegal,
				Literal: []rune{l.ch},
				Line:    l.line,
				Column:  l.column,
			})
			tok = l.newSingleToken(token.Illegal)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) newToken(tokenType token.TokenType, tokenLiteral []rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: tokenLiteral,
		Line:    l.line,
		Column:  l.column - len(tokenLiteral),
	}
}

func (l *Lexer) newSingleToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: []rune{l.ch},
		Line:    l.line,
		Column:  l.column,
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	l.column += 1
}

func (l *Lexer) consumeIdent() []rune {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) consumeString() []rune {
	pos := l.position + 1
	for {
		l.readChar()

		if l.ch == 0 || l.ch == '\n' || l.ch == ';' {
			l.Errors.AddLexerError("thiếu dấu \" kết thúc chuỗi", token.Token{
				Line:   l.line,
				Column: l.column,
			})
			l.line++
			l.column = 0
			break

		} else if l.ch == '"' {
			break
		}
	}
	return l.input[pos:l.position]
}

func (l *Lexer) consumeNumber() []rune {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) consumeSpace() []rune {
	spaces := []rune{}
	for l.peekChar() == ' ' || l.peekChar() == '\t' {
		if l.peekChar() == ' ' {
			spaces = append(spaces, ' ')
		} else if l.peekChar() == '\t' {
			spaces = append(spaces, []rune("    ")...)
		}
		l.readChar()
	}
	return spaces
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}
