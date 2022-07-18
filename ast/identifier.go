package ast

import "vila/token"

type Identifier struct {
	Tok   token.Token
	Value []rune
}

func (i *Identifier) Token() token.Token {
	return i.Tok
}

func (i *Identifier) String() string {
	return string(i.Value)
}
