package lexer

import (
	"vila/error"
	"vila/token"
)

var VNALPHA = arrToMap([]rune("aAàÀảẢãÃáÁạẠăĂằẰẳẲẵẴắẮặẶâÂầẦẩẨẫẪấẤậẬbBcCdDđĐeEèÈẻẺẽẼéÉẹẸêÊềỀểỂễỄếẾệỆfFgGhHiIìÌỉỈĩĨíÍịỊjJkKlLmMnNoOòÒỏỎõÕóÓọỌôÔồỒổỔỗỖốỐộỘơƠờỜởỞỡỠớỚợỢpPqQrRsStTuUùÙủỦũŨúÚụỤưƯừỪửỬữỮứỨựỰvVwWxXyYỳỲỷỶỹỸýÝỵỴzZ"))

type Lexer struct {
	input           []rune
	position        int
	readPosition    int
	ch              rune
	isPreviousIdent bool
	tempNextToken   token.Token
	line            int
	row             int
}

func New(input string) *Lexer {
	l := &Lexer{
		input:           []rune(input),
		readPosition:    0,
		line:            1,
		row:             0,
		isPreviousIdent: false,
		tempNextToken:   token.Token{},
	}

	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	if l.tempNextToken.Type != "" {
		nextToken := l.tempNextToken
		l.tempNextToken = token.Token{}
		return nextToken
	}

	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		tok = l.newDoubleToken(token.EQ, token.ASSIGN, '=')
	case '>':
		tok = l.newDoubleToken(token.GREATER_EQ, token.GREATER, '=')
	case '<':
		tok = l.newDoubleToken(token.LESS_EQ, token.LESS, '=')
	case '!':
		tok = l.newDoubleToken(token.NEQ, token.BANG, '=')
	case '+':
		tok = l.newSingleToken(token.PLUS)
	case '-':
		tok = l.newSingleToken(token.MINUS)
	case '*':
		tok = l.newSingleToken(token.MULTIPLY)
	case '/':
		tok = l.newSingleToken(token.SLASH)
	case '^':
		tok = l.newSingleToken(token.HAT)
	case '(':
		tok = l.newSingleToken(token.LPAREN)
	case ')':
		tok = l.newSingleToken(token.RPAREN)
	case '{':
		tok = l.newSingleToken(token.LBRACE)
	case '}':
		tok = l.newSingleToken(token.RBRACE)
	case ';':
		tok = l.newSingleToken(token.SEMICOLON)
	case '"':
		tok = l.newToken(token.STRING, l.readString())
	case '\n':
		tok = l.newSingleToken(token.ENDLINE)
		l.line += 1
		l.row = 0
	case 0:
		tok = l.newToken(token.EOF, []rune{})
	default:
		if isLetter(l.ch) {
			tokenLiteral := l.readWord()
			tokenType := token.LookupKeyword(tokenLiteral)
			tok = l.newToken(tokenType, tokenLiteral)

			// Handle spacing identifier
			if tokenType == token.IDENT && !l.isPreviousIdent {

				l.isPreviousIdent = true
				nextToken := l.NextToken()

				for nextToken.Type == token.IDENT {
					appendIdent(&tok, nextToken)
					nextToken = l.NextToken()
				}

				l.isPreviousIdent = false
				l.tempNextToken = nextToken
			}

			return tok

		} else if isDigit(l.ch) {
			tok = l.newToken(token.INT, l.readNumber())
			return tok
		} else {
			tok = l.newSingleToken(token.ILLEGAL)
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
		Row:     l.row - len(tokenLiteral),
	}
}

func (l *Lexer) newDoubleToken(tokenType token.TokenType, defaultTokenType token.TokenType, nextChar rune) token.Token {
	if l.peekChar() == nextChar {
		eq := l.ch
		l.readChar()
		return l.newToken(tokenType, []rune{eq, l.ch})
	} else {
		return l.newSingleToken(defaultTokenType)
	}
}

func (l *Lexer) newSingleToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: []rune{l.ch},
		Line:    l.line,
		Row:     l.row,
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
	l.row += 1
}

func (l *Lexer) readWord() []rune {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readString() []rune {
	pos := l.position + 1
	for {
		l.readChar()

		if l.ch == 0 {
			error.ThrowError(error.SYNTAX_ERROR, "thiếu dấu \" kết thúc chuỗi")

		} else if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() []rune {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
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
