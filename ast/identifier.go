package ast

import "vila/token"

type Identifier struct {
	Token token.Token
	Value []rune
}

func (i *Identifier) String() string {
	return string(i.Value)
}
