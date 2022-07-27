package ast

import "vila/token"

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func (ie *IfExpression) FromToken() token.Token {
	return ie.Token
}

func (ie *IfExpression) ToToken() token.Token {
	return ie.Alternative.ToToken()
}

func (is *IfExpression) String() string {
	return ""
}
