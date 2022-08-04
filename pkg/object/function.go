package object

import (
	"bytes"
	"vila/pkg/ast"
)

const (
	FUNC_OBJ = "Hàm"
)

type Function struct {
	Ident  *ast.Identifier
	Params []*ast.Identifier
	Body   ast.Expression
	Env    *Environment

	Builtin     func(args ...Object) Object
	LeftCompose *Function
}

func (fn *Function) Type() ObjectType { return FUNC_OBJ }
func (fn *Function) Display() string {
	if fn.Builtin != nil {
		return "<Hàm cài đặt sẵn>"
	}

	var out bytes.Buffer

	if fn.LeftCompose != nil {
		out.WriteString("(")
	}

	out.WriteString(fn.Ident.Value)

	tempFn := fn
	for tempFn.LeftCompose != nil {
		tempFn = tempFn.LeftCompose
		out.WriteString("." + tempFn.Ident.Value)
	}

	if fn.LeftCompose != nil {
		out.WriteString(")")
	}

	out.WriteString("(")
	for i, param := range fn.Params {
		out.WriteString(param.Value)
		if i != len(fn.Params)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString(")")

	return out.String()
}

func (fn *Function) Dot(right Object) Object {
	switch right := right.(type) {
	case *Function:
		newFn := *right
		newFn.LeftCompose = fn
		return &newFn
	default:
		return CANT_OPERATE
	}
}
