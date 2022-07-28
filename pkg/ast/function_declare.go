package ast

import (
	"vila/pkg/token"
)

type FunctionDeclareStatement struct {
	Token  token.Token
	Ident  *Identifier
	Params []*Identifier
	Body   Expression
}

func (fn *FunctionDeclareStatement) FromToken() token.Token {
	return fn.Token
}

func (fn *FunctionDeclareStatement) ToToken() token.Token {
	return fn.Body.ToToken()
}

func (fn *FunctionDeclareStatement) String() string {
	return ""
}
