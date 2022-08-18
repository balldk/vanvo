package ast

import (
	"bytes"
	"vanvo/pkg/token"
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

func (ib *IfBranch) String() string {
	var out bytes.Buffer

	out.WriteString(string(ib.Token.Literal) + " ")
	out.WriteString(ib.Condition.String())
	out.WriteString(ib.Consequence.String())

	return out.String()
}

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
	var out bytes.Buffer

	for _, branch := range is.Branches {
		out.WriteString(branch.String() + "\n")
	}

	return out.String()
}
