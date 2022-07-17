package ast

import "bytes"

type RealInterval struct {
	Lower Expression
	Upper Expression
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
	Lower Expression
	Upper Expression
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
