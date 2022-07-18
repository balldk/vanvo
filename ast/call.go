package ast

import (
	"bytes"
	"strings"
	"vila/token"
)

type CallExpression struct {
	Tok       token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) Token() token.Token {
	return ce.Tok
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
