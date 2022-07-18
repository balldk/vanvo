package ast

import (
	"bytes"
	"vila/token"
)

type AssignStatement struct {
	Tok   token.Token
	Ident *Identifier
	Value Expression
}

func (as *AssignStatement) Token() token.Token {
	return as.Tok
}

func (as *AssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.Ident.String())
	out.WriteString(" = ")
	out.WriteString(as.Value.String())

	return out.String()
}
