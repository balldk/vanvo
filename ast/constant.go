package ast

import "vila/token"

type Int struct {
	Token token.Token
	Value int64
}

func (i *Int) FromToken() token.Token { return i.Token }
func (i *Int) ToToken() token.Token   { return i.Token }
func (i *Int) String() string {
	return string(i.Token.Literal)
}

type Real struct {
	Token token.Token
	Value float64
}

func (r *Real) FromToken() token.Token {
	return r.Token
}

func (r *Real) ToToken() token.Token {
	return r.Token
}

func (i *Real) String() string {
	return string(i.Token.Literal)
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) FromToken() token.Token { return b.Token }
func (b *Boolean) ToToken() token.Token   { return b.Token }
func (b *Boolean) String() string {
	return string(b.Token.Literal)
}
