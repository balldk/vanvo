package ast

import "vila/token"

type Int struct {
	Token token.Token
	Value int64
}

func (i *Int) String() string {
	return string(i.Token.Literal)
}

type Real struct {
	Token token.Token
	Value float64
}

func (i *Real) String() string {
	return string(i.Token.Literal)
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) String() string {
	return string(b.Token.Literal)
}
