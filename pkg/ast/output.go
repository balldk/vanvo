package ast

import (
	"bytes"
	"vila/pkg/token"
)

type OutputStatement struct {
	Token  token.Token
	Values []Expression
}

func (os *OutputStatement) FromToken() token.Token {
	return os.Token
}

func (os *OutputStatement) ToToken() token.Token {
	if len(os.Values) > 0 {
		return os.Values[len(os.Values)-1].ToToken()
	}
	return token.Token{}
}

func (os *OutputStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(os.Token.Literal) + " ")

	for _, value := range os.Values {
		out.WriteString(value.String())
		out.WriteString(", ")
	}

	return out.String()
}
