package ast

import (
	"bytes"
	"vila/token"
)

type BlockStatement struct {
	LeftBrace  token.Token
	RightBrace token.Token
	Statements []Statement
}

func (bs *BlockStatement) FromToken() token.Token {
	return bs.LeftBrace
}

func (bs *BlockStatement) ToToken() token.Token {
	return bs.RightBrace
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
