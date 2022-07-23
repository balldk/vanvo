package ast

import (
	"bytes"
	"vila/token"
)

type PrefixExpression struct {
	Operator token.Token
	Right    Expression
}

func (pe *PrefixExpression) FromToken() token.Token {
	return pe.Operator
}

func (pe *PrefixExpression) ToToken() token.Token {
	return pe.Right.ToToken()
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(Prefix: ")
	out.WriteString(string(pe.Operator.Literal))
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (ie *InfixExpression) FromToken() token.Token {
	return ie.Left.FromToken()
}

func (ie *InfixExpression) ToToken() token.Token {
	return ie.Right.ToToken()
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("{ Operator: " + string(ie.Operator.Literal) + "\n")
	out.WriteString("Left: " + ie.Left.String() + "\n")
	out.WriteString("Right: " + ie.Right.String() + " }")

	return out.String()
}
