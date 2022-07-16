package token

import (
	"fmt"
)

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	ENDLINE = "\n"

	IDENT  = "IDENT"
	INT    = "INT"
	REAL   = "REAL"
	STRING = "STRING"
	TRUE   = "TRUE"
	FALSE  = "FALSE"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	PERCENT  = "%"
	SLASH    = "/"
	HAT      = "^"
	BANG     = "!"
	DOT      = "."

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
	COLON     = "."
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
