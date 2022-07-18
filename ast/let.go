package ast

import (
	"bytes"
	"vila/token"
)

type LetStatement struct {
	Tok     token.Token
	Ident   *Identifier
	Value   Expression
	SetType Expression
}

func (ls *LetStatement) Token() token.Token {
	return ls.Tok
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(ls.Tok.Literal) + " ")
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
