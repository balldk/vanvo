package ast

import "vila/token"

type Node interface {
	String() string
	Token() token.Token
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type ExpressionStatement struct {
	Tok        token.Token
	Expression Expression
}

func (r *ExpressionStatement) Token() token.Token {
	return r.Tok
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
