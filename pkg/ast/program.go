package ast

import (
	"bytes"
	"vila/pkg/token"
)

type Program struct {
	Statements []Statement
}

func (p *Program) FromToken() token.Token {
	if len(p.Statements) > 0 {
		return p.Statements[0].FromToken()
	}
	return token.Token{}
}

func (p *Program) ToToken() token.Token {
	if len(p.Statements) > 0 {
		return p.Statements[len(p.Statements)-1].ToToken()
	}
	return token.Token{}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String() + "\n")
	}

	return out.String()
}
