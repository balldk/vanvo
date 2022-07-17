package token

import (
	"fmt"
)

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	ENDLINE = "ENDLINE"

	IDENT  = "IDENT"
	INT    = "Số nguyên"
	REAL   = "Số thực"
	STRING = "Chuỗi"
	TRUE   = "đúng"
	FALSE  = "sai"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	PERCENT  = "%"
	SLASH    = "/"
	HAT      = "^"
	BANG     = "!"
	DOT      = "."
	DOTDOT   = ".."

	EQ         = "=="
	NEQ        = "!="
	LESS       = "<"
	GREATER    = ">"
	LESS_EQ    = "<="
	GREATER_EQ = ">="

	LET     = "cho"
	IF      = "nếu"
	ELSE_IF = "còn nếu"
	ELSE    = "còn không"
	FOR     = "với"
	BELONG  = "thuộc"
	IMPLY   = "=>"
	INPUT   = "nhập"
	OUTPUT  = "xuất"

	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
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
	"cho":       LET,
	"nếu":       IF,
	"còn nếu":   ELSE_IF,
	"còn không": ELSE,
	"đúng":      TRUE,
	"sai":       FALSE,
	"với":       FOR,
	"thuộc":     BELONG,
	"nhập":      INPUT,
	"xuất":      OUTPUT,
})

func LookupKeyword(word []rune) TokenType {
	if tok, ok := keywords[string(word)]; ok {
		return tok
	}
	return IDENT
}
