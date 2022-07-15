package ast

import "vila/token"

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) String() string {
	return string(b.Token.Literal)
}
