package ast

import (
	"bytes"
	"vila/pkg/token"
)

type BlockStatement struct {
	Colon      token.Token
	Statements []Statement
}

func (bs *BlockStatement) FromToken() token.Token {
	return bs.Colon
}

func (bs *BlockStatement) ToToken() token.Token {
	if len(bs.Statements) > 0 {
		return bs.Statements[len(bs.Statements)-1].ToToken()
	}
	return bs.Colon
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString(":\n")
	for _, s := range bs.Statements {
		out.WriteString(s.String() + "\n")
	}
	return out.String()
}
