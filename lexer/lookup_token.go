package lexer

import (
	"vila/token"
)

var TOKEN_TABLE = map[string]token.TokenType{
	"":   token.EOF,
	"=":  token.ASSIGN,
	"!":  token.BANG,
	"+":  token.PLUS,
	"-":  token.MINUS,
	"*":  token.ASTERISK,
	"/":  token.SLASH,
	"^":  token.HAT,
	">":  token.GREATER,
	"<":  token.EQ,
	"(":  token.LPAREN,
	")":  token.RPAREN,
	"{":  token.LBRACE,
	"}":  token.RBRACE,
	"[":  token.LBRACKET,
	"]":  token.RBRACKET,
	".":  token.DOT,
	",":  token.COMMA,
	":":  token.COLON,
	";":  token.SEMICOLON,
	"==": token.EQ,
	"<=": token.LESS_EQ,
	">=": token.GREATER_EQ,
	"=>": token.IMPLY,
}

func (l *Lexer) lookupToken() token.Token {
	doubleCh := string([]rune{l.ch, l.peekChar()})
	ch := string(l.ch)

	if tokenType, ok := TOKEN_TABLE[doubleCh]; ok {
		l.readChar()
		return token.Token{
			Type:    tokenType,
			Literal: []rune(doubleCh),
			Line:    l.line,
			Row:     l.row,
		}
	} else if tokenType, ok := TOKEN_TABLE[ch]; ok {
		return token.Token{
			Type:    tokenType,
			Literal: []rune(ch),
			Line:    l.line,
			Row:     l.row,
		}
	}

	return token.Token{Type: token.ILLEGAL}
}
