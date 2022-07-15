package ast

import "vila/token"

type Real struct {
	Token token.Token
	Value float64
}

func (i *Real) String() string {
	return string(i.Token.Literal)
}
