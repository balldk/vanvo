package ast

import (
	"bytes"
	"vila/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) FromToken() token.Token {
	return ie.Token
}

func (ie *IfExpression) ToToken() token.Token {
	if ie.Alternative != nil {
		return ie.Alternative.ToToken()
	}
	return ie.Consequence.ToToken()
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
