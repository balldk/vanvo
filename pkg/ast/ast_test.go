package ast

import (
	"testing"
	"vanvo/pkg/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarDeclareStatement{
				Token: token.Token{Type: token.Let, Literal: []rune("cho")},
				Ident: &Identifier{
					Token: token.Token{Type: token.Ident, Literal: []rune("myName")},
					Value: "myName",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.Ident, Literal: []rune("otherName")},
					Value: "otherName",
				},
			},
		},
	}

	if program.String() != "cho myName = otherName\n" {
		t.Errorf("program.String() wrong. Got=%q", program.String())
	}
}
