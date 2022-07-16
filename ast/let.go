package ast

import (
	"bytes"
	"vila/token"
)

type LetStatement struct {
	Token   token.Token
	Ident   *Identifier
	Value   Expression
	SetType Expression
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(ls.Token.Literal) + " ")
	out.WriteString(ls.Ident.String())

	if ls.Value != nil {
		out.WriteString(" = ")
		out.WriteString(ls.Value.String())
	}

	if ls.SetType != nil {
		out.WriteString(" thuộc ")
		out.WriteString(ls.SetType.String())
	}

	return out.String()
}
