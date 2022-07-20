package token

import (
	"fmt"
)

type TokenType string

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"
	Endline = "Endline"

	Ident  = "IDENT"
	Int    = "Số nguyên"
	Real   = "Số thực"
	String = "Chuỗi"
	True   = "đúng"
	False  = "sai"

	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Asterisk = "*"
	Percent  = "%"
	Slash    = "/"
	Hat      = "^"
	Bang     = "!"
	Dot      = "."
	DotDot   = ".."

	Equal        = "=="
	NotEqual     = "!="
	Less         = "<"
	Greater      = ">"
	LessEqual    = "<="
	GreaterEqual = ">="

	Let    = "cho"
	If     = "nếu"
	ElseIf = "còn nếu"
	Else   = "còn không"
	For    = "với"
	Belong = "thuộc"
	Imply  = "=>"
	Input  = "nhập"
	Output = "xuất"

	LParen    = "("
	RParen    = ")"
	LBrace    = "{"
	RBrace    = "}"
	LBracket  = "["
	RBracket  = "]"
	Comma     = ","
	Colon     = ":"
	Semicolon = ";"
)

type Token struct {
	Type    TokenType
	Line    int
	Row     int
	Literal []rune
}

func (t Token) String() string {
	return fmt.Sprintf("{ Literal: %s, Type: %v, Line: %d, Row: %d }", string(t.Literal), t.Type, t.Line, t.Row)
}

var keywords = keywordsWithoutDiacritic(map[string]TokenType{
	"cho":       Let,
	"nếu":       If,
	"còn nếu":   ElseIf,
	"còn không": Else,
	"đúng":      True,
	"sai":       False,
	"với":       For,
	"thuộc":     Belong,
	"nhập":      Input,
	"xuất":      Output,
})

func LookupKeyword(word []rune) TokenType {
	if tok, ok := keywords[string(word)]; ok {
		return tok
	}
	return Ident
}
