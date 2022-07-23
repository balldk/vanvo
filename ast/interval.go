package ast

import (
	"bytes"
	"vila/token"
)

type RealInterval struct {
	LeftBracket  token.Token
	RightBracket token.Token
	Lower        Expression
	Upper        Expression
}

func (r *RealInterval) FromToken() token.Token {
	return r.LeftBracket
}

func (r *RealInterval) ToToken() token.Token {
	return r.RightBracket
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
	LeftBracket  token.Token
	RightBracket token.Token
	Lower        Expression
	Upper        Expression
}

func (ii *IntInterval) FromToken() token.Token {
	return ii.LeftBracket
}

func (ii *IntInterval) ToToken() token.Token {
	return ii.RightBracket
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
