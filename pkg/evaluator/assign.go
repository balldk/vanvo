package evaluator

import (
	"fmt"
	"vila/pkg/ast"
	"vila/pkg/object"
)

func (ev *Evaluator) evalVarDeclare(node *ast.VarDeclareStatement) object.Object {

	val := ev.Eval(node.Value)
	if val == NULL {
		return NULL
	}

	if _, ok := ev.Env.GetInScope(node.Ident.Value); ok {
		errMsg := fmt.Sprintf("'%s' đã được khởi tạo", node.Ident.Value)
		ev.runtimeError(errMsg)
	}
	ev.Env.SetInScope(node.Ident.Value, val)

	return val
}

func (ev *Evaluator) evalFunctionDeclare(node *ast.FunctionDeclareStatement) {

	params := node.Params
	body := node.Body
	fn := &object.Function{Ident: node.Ident, Params: params, Body: body}
	if _, ok := ev.Env.GetInScope(node.Ident.Value); ok {
		errMsg := fmt.Sprintf("'%s' đã được khởi tạo", node.Ident.Value)
		ev.runtimeError(errMsg)
	}
	ev.Env.SetInScope(node.Ident.Value, fn)
}

func (ev *Evaluator) evalAssignStatement(node *ast.AssignStatement) object.Object {

	if _, ok := object.Builtins[node.Ident.Value]; ok {
		errMsg := fmt.Sprintf("Không thể gán giá trị cho '%s'", node.Ident.Value)
		return ev.runtimeError(errMsg)
	}

	val := ev.Eval(node.Value)
	if val == NULL {
		return NULL
	}

	obj := ev.Env.Set(node.Ident.Value, val)
	if obj == nil {
		errMsg := fmt.Sprintf("'%s' chưa được khai tạo", node.Ident.Value)
		return ev.runtimeError(errMsg)
	}

	return val
}
