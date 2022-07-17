package ast

import (
	"bytes"
	"vila/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator []rune
	Right    Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(string(pe.Operator))
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

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
