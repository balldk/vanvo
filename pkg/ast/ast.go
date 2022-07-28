package ast

import (
	"vila/pkg/token"
)

type Node interface {
	String() string
	FromToken() token.Token
	ToToken() token.Token
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type ExpressionStatement struct {
	Expression Expression
}

func (r *ExpressionStatement) FromToken() token.Token {
	return r.Expression.FromToken()
}

func (r *ExpressionStatement) ToToken() token.Token {
	return r.Expression.ToToken()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
