package ast

import (
	"bytes"
	"vila/token"
)

type BlockStatement struct {
	Tok        token.Token
	Statements []Statement
}

func (bs *BlockStatement) Token() token.Token {
	return bs.Tok
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
