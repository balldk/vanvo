package ast

import (
	"vila/pkg/token"
)

type IfBranch struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (ib *IfBranch) FromToken() token.Token {
	return ib.Token
}

func (ib *IfBranch) ToToken() token.Token {
	return ib.Consequence.ToToken()
}

func (ib *IfBranch) String() string { return "" }

type IfStatement struct {
	Branches []*IfBranch
}

func (is *IfStatement) FromToken() token.Token {
	return is.Branches[0].FromToken()
}

func (is *IfStatement) ToToken() token.Token {
	return is.Branches[len(is.Branches)-1].FromToken()
}

func (is *IfStatement) String() string {
	return ""
}
