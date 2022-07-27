package ast

import "vila/token"

type FunctionDeclare struct {
	Name   *Identifier
	Params []*Identifier
	Body   *GroupExpression
}

func (fn *FunctionDeclare) FromToken() token.Token {
	return fn.Name.ToToken()
}

func (fn *FunctionDeclare) ToToken() token.Token {
	return fn.Body.ToToken()
}
