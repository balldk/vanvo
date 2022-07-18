package ast

import "vila/token"

type Int struct {
	Tok   token.Token
	Value int64
}

func (i *Int) Token() token.Token {
	return i.Tok
}

func (i *Int) String() string {
	return string(i.Tok.Literal)
}

type Real struct {
	Tok   token.Token
	Value float64
}

func (r *Real) Token() token.Token {
	return r.Tok
}

func (i *Real) String() string {
	return string(i.Tok.Literal)
}

type Boolean struct {
	Tok   token.Token
	Value bool
}

func (b *Boolean) Token() token.Token {
	return b.Tok
}

func (b *Boolean) String() string {
	return string(b.Tok.Literal)
}
