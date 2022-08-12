package ast

import (
	"bytes"
	"vila/pkg/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func (ie *IfExpression) FromToken() token.Token {
	return ie.Token
}

func (ie *IfExpression) ToToken() token.Token {
	return ie.Alternative.ToToken()
}

func (is *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString(is.Consequence.String())
	out.WriteString(" nếu ")
	out.WriteString(is.Condition.String())
	out.WriteString(" còn không ")
	out.WriteString(is.Alternative.String())

	return out.String()
}
