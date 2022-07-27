package ast

import (
	"bytes"
	"vila/token"
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
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
