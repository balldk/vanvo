package ast

import "vila/token"

type Int struct {
	Token token.Token
	Value int64
}

func (i *Int) String() string {
	return string(i.Token.Literal)
}
