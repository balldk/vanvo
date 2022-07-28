package ast

import (
	"bytes"
	"strings"
	"vila/pkg/token"
)

type CallExpression struct {
	RightParen token.Token
	Function   Expression
	Arguments  []Expression
}

func (ce *CallExpression) FromToken() token.Token {
	return ce.Function.FromToken()
}

func (ce *CallExpression) ToToken() token.Token {
	return ce.RightParen
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
