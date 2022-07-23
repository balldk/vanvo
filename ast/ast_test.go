package ast

import (
	"testing"
	"vila/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.Let, Literal: []rune("cho")},
				Ident: &Identifier{
					Token: token.Token{Type: token.Ident, Literal: []rune("myName")},
					Value: []rune("myName"),
				},
				Value: &Identifier{
					Token: token.Token{Type: token.Ident, Literal: []rune("otherName")},
					Value: []rune("otherName"),
				},
			},
		},
	}

	if program.String() != "cho myName = otherName\n" {
		t.Errorf("program.String() wrong. Got=%q", program.String())
	}
}
