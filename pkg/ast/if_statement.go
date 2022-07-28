package ast

import (
	"bytes"
	"vila/pkg/token"
)

type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) FromToken() token.Token {
	return is.Token
}

func (is *IfStatement) ToToken() token.Token {
	if is.Alternative != nil {
		return is.Alternative.ToToken()
	}
	return is.Consequence.ToToken()
}

func (is *IfStatement) String() string {
	var out bytes.Buffer

	out.WriteString("nếu ")
	out.WriteString(is.Condition.String())
	out.WriteString(" ")
	out.WriteString(is.Consequence.String())

	if is.Alternative != nil {
		out.WriteString("còn không ")
		out.WriteString(is.Alternative.String())
	}

	return out.String()
}
