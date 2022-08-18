package evaluator

import (
	"fmt"
	"vanvo/pkg/ast"
	"vanvo/pkg/object"
)

func (ev *Evaluator) evalCallExpression(call *ast.CallExpression) object.Object {
	fn := ev.Eval(call.Function)
	args := ev.evalExpressions(call.Arguments)

	switch fn := fn.(type) {
	case *object.Function:
		res := ev.applyFunction(fn, args)
		args = []object.Object{res}

		for fn.LeftCompose != nil {
			fn = fn.LeftCompose
			res = ev.applyFunction(fn, args)
			args = []object.Object{res}
		}

		return res

	default:
		if len(call.Arguments) == 1 {
			right := ev.Eval(call.Arguments[0])
			return ev.evalMultiplication(fn, right)
		}

		return ev.runtimeError("Biểu thức không hợp lệ")
	}
}

func (ev *Evaluator) applyFunction(fn *object.Function, args []object.Object) object.Object {
	env := object.NewEnclosedEnvironment(ev.Env)

	if fn.Builtin != nil {
		res := fn.Builtin(args...)
		if err, ok := res.(*object.Error); ok {
			return ev.runtimeError(err.Message)
		}
		if err, ok := res.(*object.ArgumentError); ok {
			errMsg := fmt.Sprintf("Cần %d tham số thay vì %d", err.Expected, err.Received)
			return ev.runtimeError(errMsg)
		}

		return res
	}

	if len(args) != len(fn.Params) {
		errMsg := fmt.Sprintf(
			"'%s' cần %d tham số thay vì %d",
			fn.Ident.Value, len(fn.Params), len(args))

		return ev.runtimeError(errMsg)
	}

	for index, param := range fn.Params {
		env.SetInScope(param.Value, args[index])
	}

	val := ev.Eval(fn.Body, env)
	return ev.unwrapImply(val)
}

func (ev *Evaluator) unwrapImply(obj object.Object) object.Object {
	if imply, ok := obj.(*object.Imply); ok {
		return ev.unwrapImply(imply.Value)
	}

	return obj
}
