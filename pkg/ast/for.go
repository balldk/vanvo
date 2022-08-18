package ast

import "vanvo/pkg/token"

type ForStatement struct {
	Token      token.Token
	Conditions []Expression
	Body       *BlockStatement
}

func (fs *ForStatement) FromToken() token.Token {
	return fs.Token
}

func (fs *ForStatement) ToToken() token.Token {
	return fs.Body.ToToken()
}

func (fs *ForStatement) String() string { return "" }
