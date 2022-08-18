package ast

import "vanvo/pkg/token"

type String struct {
	Token token.Token
	Value string
}

func (s *String) FromToken() token.Token { return s.Token }
func (s *String) ToToken() token.Token   { return s.Token }
func (s *String) String() string {
	return s.Value
}
