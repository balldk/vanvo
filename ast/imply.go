package ast

import (
	"bytes"
	"vila/token"
)

type ImplyStatement struct {
	Tok   token.Token
	Value Expression
}

func (is *ImplyStatement) Token() token.Token {
	return is.Tok
}

func (is *ImplyStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(is.Tok.Literal) + " ")

	if is.Value != nil {
		out.WriteString(is.Value.String())
	}

	return out.String()
}
