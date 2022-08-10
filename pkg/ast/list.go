package ast

import (
	"vila/pkg/token"
)

type List struct {
	LeftBrace  token.Token
	RightBrace token.Token
	Data       []Expression
}

func (li *List) FromToken() token.Token {
	return li.LeftBrace
}

func (li *List) ToToken() token.Token {
	return li.RightBrace
}

func (li *List) String() string {
	return ""
}
