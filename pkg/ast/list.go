package ast

import (
	"vanvo/pkg/token"
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

type ListComprehension struct {
	LeftBrace  token.Token
	RightBrace token.Token
	Expression Expression
	Conditions []Expression
}

func (li *ListComprehension) FromToken() token.Token {
	return li.LeftBrace
}

func (li *ListComprehension) ToToken() token.Token {
	return li.RightBrace
}

func (li *ListComprehension) String() string {
	return ""
}
