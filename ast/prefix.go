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
