package ast

import (
	"bytes"
	"vanvo/pkg/token"
)

type VarDeclareStatement struct {
	Token token.Token
	Ident *Identifier
	Value Expression
}

func (vds *VarDeclareStatement) FromToken() token.Token {
	return vds.Token
}

func (vds *VarDeclareStatement) ToToken() token.Token {
	return vds.Value.ToToken()
}

func (vds *VarDeclareStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(vds.Token.Literal) + " ")
	out.WriteString(vds.Ident.String())

	if vds.Value != nil {
		out.WriteString(" = ")
		out.WriteString(vds.Value.String())
	}

	return out.String()
}
