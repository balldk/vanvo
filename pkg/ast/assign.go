package ast

import (
	"bytes"
	"vila/pkg/token"
)

type AssignStatement struct {
	Token token.Token
	Ident *Identifier
	Value Expression
}

func (as *AssignStatement) FromToken() token.Token {
	return as.Ident.FromToken()
}

func (as *AssignStatement) ToToken() token.Token {
	return as.Value.ToToken()
}

func (as *AssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.Ident.String())
	out.WriteString(" = ")
	out.WriteString(as.Value.String())

	return out.String()
}
