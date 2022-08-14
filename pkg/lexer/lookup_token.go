package lexer

import (
	"vila/pkg/token"
)

var TOKEN_TABLE = map[string]token.TokenType{
	"":   token.EOF,
	"=":  token.Assign,
	"!":  token.Bang,
	"#":  token.Hash,
	"+":  token.Plus,
	"-":  token.Minus,
	"*":  token.Asterisk,
	"/":  token.Slash,
	"%":  token.Percent,
	"^":  token.Hat,
	">":  token.Greater,
	"<":  token.Less,
	"(":  token.LParen,
	")":  token.RParen,
	"{":  token.LBrace,
	"}":  token.RBrace,
	"[":  token.LBracket,
	"]":  token.RBracket,
	".":  token.Dot,
	"..": token.DotDot,
	",":  token.Comma,
	":":  token.Colon,
	";":  token.Semicolon,
	"==": token.Equal,
	"!=": token.NotEqual,
	"<=": token.LessEqual,
	">=": token.GreaterEqual,
	"=>": token.Imply,
	"//": token.SlashSlash,
	"|":  token.Bar,
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
			Column:  l.column,
		}
	} else if tokenType, ok := TOKEN_TABLE[ch]; ok {
		return token.Token{
			Type:    tokenType,
			Literal: []rune(ch),
			Line:    l.line,
			Column:  l.column,
		}
	}

	return token.Token{Type: token.Illegal}
}
