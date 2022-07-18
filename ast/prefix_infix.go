package ast

import (
	"bytes"
	"vila/token"
)

type PrefixExpression struct {
	Operator token.Token
	Right    Expression
}

func (pe *PrefixExpression) Token() token.Token {
	return pe.Operator
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(string(pe.Operator.Literal))
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (ie *InfixExpression) Token() token.Token {
	return ie.Operator
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + string(ie.Operator.Literal) + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
