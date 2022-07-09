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
	INT    = "INT"
	STRING = "STRING"
	TRUE   = "TRUE"
	FALSE  = "FALSE"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	SLASH    = "/"
	HAT      = "^"
	BANG     = "!"

	EQ         = "=="
	NEQ        = "!="
	LESS       = "<"
	GREATER    = ">"
	LESS_EQ    = "<="
	GREATER_EQ = ">="

	LET    = "LET"
	IF     = "IF"
	FOR    = "FOR"
	IN     = "IN"
	INPUT  = "INPUT"
	OUTPUT = "OUTPUT"

	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"
	SEMICOLON = "SEMICOLON"
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
	"cho":   LET,
	"nếu":   IF,
	"đúng":  TRUE,
	"sai":   FALSE,
	"với":   FOR,
	"thuộc": IN,
	"nhập":  INPUT,
	"xuất":  OUTPUT,
})

func LookupKeyword(word []rune) TokenType {
	if tok, ok := keywords[string(word)]; ok {
		return tok
	}
	return IDENT
}
