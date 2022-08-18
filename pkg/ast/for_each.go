package ast

import "vanvo/pkg/token"

type ForEachStatement struct {
	Token      token.Token
	Conditions []Expression
	Body       *BlockStatement
}

func (fe *ForEachStatement) FromToken() token.Token {
	return fe.Token
}

func (fe *ForEachStatement) ToToken() token.Token {
	return fe.Body.ToToken()
}

func (fe *ForEachStatement) String() string { return "" }
