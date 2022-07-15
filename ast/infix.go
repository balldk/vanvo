package ast

import (
	"bytes"
	"vila/token"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator []rune
	Right    Expression
}

func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + string(oe.Operator) + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}
