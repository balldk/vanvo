package object

import (
	"bytes"
	"strings"
	"vila/pkg/ast"
)

const (
	FUNC_OBJ = "Hàm"
)

type Function struct {
	Ident  ast.Identifier
	Params []*ast.Identifier
	Body   ast.Expression
	Env    *Environment
}

func (fn *Function) Type() ObjectType { return FUNC_OBJ }
func (fn *Function) Display() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fn.Params {
		params = append(params, p.String())
	}

	out.WriteString(fn.Ident.Value)
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") = ")
	out.WriteString(fn.Body.String())

	return out.String()
}
