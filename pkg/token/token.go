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

	Equal        = "=="
	NotEqual     = "!="
	Less         = "<"
	Greater      = ">"
	LessEqual    = "<="
	GreaterEqual = ">="

	Let     = "cho"
	If      = "nếu"
	ElseIf  = "còn nếu"
	Else    = "còn không"
	For     = "với"
	ForEach = "với mỗi"
	Belong  = "thuộc"
	Imply   = "=>"
	Input   = "nhập"
	Output  = "xuất"

	LParen     = "("
	RParen     = ")"
	LBrace     = "{"
	RBrace     = "}"
	LBracket   = "["
	RBracket   = "]"
	Comma      = ","
	Colon      = ":"
	Semicolon  = ";"
	DotDot     = ".."
	SlashSlash = "//"
)

type Token struct {
	Type    TokenType
	Line    int
	Column  int
	Literal []rune
}

func (t Token) String() string {
	return fmt.Sprintf("{ Literal: %s, Type: %v, Line: %d, Column: %d }", string(t.Literal), t.Type, t.Line, t.Column)
}

var keywords = keywordsWithoutDiacritic(map[string]TokenType{
	"cho":       Let,
	"nếu":       If,
	"còn nếu":   ElseIf,
	"còn không": Else,
	"đúng":      True,
	"sai":       False,
	"với":       For,
	"với mỗi":   ForEach,
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
