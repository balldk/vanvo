package ast

import (
	"bytes"
	"vila/pkg/token"
)

type IndexExpression struct {
	RightBracket token.Token
	Set          Expression
	Index        Expression
}

func (ie *IndexExpression) FromToken() token.Token {
	return ie.Set.FromToken()
}

func (ie *IndexExpression) ToToken() token.Token {
	return ie.RightBracket
}

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ie.Set.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("]")

	return out.String()
}
