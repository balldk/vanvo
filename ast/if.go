package ast

import (
	"bytes"
	"vila/token"
)

type IfExpression struct {
	Tok         token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) Token() token.Token {
	return ie.Tok
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("nếu ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("còn không ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
