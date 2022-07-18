package ast

import (
	"bytes"
	"vila/token"
)

type RealInterval struct {
	Tok   token.Token
	Lower Expression
	Upper Expression
}

func (r *RealInterval) Token() token.Token {
	return r.Tok
}

func (ri *RealInterval) String() string {
	var out bytes.Buffer

	out.WriteString("[")
	out.WriteString(ri.Lower.String())
	out.WriteString(",")
	out.WriteString(ri.Upper.String())
	out.WriteString("]")

	return out.String()
}

type IntInterval struct {
	Tok   token.Token
	Lower Expression
	Upper Expression
}

func (ii *IntInterval) Token() token.Token {
	return ii.Tok
}

func (ii *IntInterval) String() string {
	var out bytes.Buffer

	out.WriteString("[")
	out.WriteString(ii.Lower.String())
	out.WriteString("..")
	out.WriteString(ii.Upper.String())
	out.WriteString("]")

	return out.String()
}
