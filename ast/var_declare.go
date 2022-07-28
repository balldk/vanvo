package ast

import (
	"bytes"
	"vila/token"
)

type VarDeclareStatement struct {
	Token   token.Token
	Ident   *Identifier
	Value   Expression
	SetType Expression
}

func (vds *VarDeclareStatement) FromToken() token.Token {
	return vds.Token
}

func (vds *VarDeclareStatement) ToToken() token.Token {
	if vds.Value == nil {
		return vds.SetType.ToToken()
	}
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

	if vds.SetType != nil {
		out.WriteString(" thuộc ")
		out.WriteString(vds.SetType.String())
	}

	return out.String()
}
