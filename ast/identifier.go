package ast

import "vila/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) FromToken() token.Token {
	return i.Token
}

func (i *Identifier) ToToken() token.Token {
	return i.Token
}

func (i *Identifier) String() string {
	return string(i.Value)
}
