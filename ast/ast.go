package ast

import "vila/token"

type Node interface {
	String() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
