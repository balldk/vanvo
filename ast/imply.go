package ast

import (
	"bytes"
	"vila/token"
)

type ImplyStatement struct {
	Token token.Token
	Value Expression
}

func (is *ImplyStatement) FromToken() token.Token {
	return is.Token
}

func (is *ImplyStatement) ToToken() token.Token {
	return is.Value.ToToken()
}

func (is *ImplyStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(is.Token.Literal) + " ")

	if is.Value != nil {
		out.WriteString(is.Value.String())
	}

	return out.String()
}
