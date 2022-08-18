package ast

import (
	"bytes"
	"vanvo/pkg/token"
)

type GroupExpression struct {
	LeftParen  token.Token
	RightParen token.Token
	Statements []Statement
}

func (bs *GroupExpression) FromToken() token.Token {
	return bs.LeftParen
}

func (bs *GroupExpression) ToToken() token.Token {
	return bs.RightParen
}

func (bs *GroupExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(\n")
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	out.WriteString(")")
	return out.String()
}
