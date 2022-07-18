package ast

import (
	"bytes"
	"vila/token"
)

type Program struct {
	Statements []Statement
}

func (p *Program) Token() token.Token {
	return token.Token{}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String() + "\n")
	}

	return out.String()
}
