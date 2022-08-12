package evaluator

import (
	"fmt"
	"vila/pkg/ast"
	"vila/pkg/object"
)

func (ev *Evaluator) evalList(list *ast.List) object.Object {
	exps := ev.evalExpressions(list.Data)
	return &object.List{Data: exps}
}

func (ev *Evaluator) evalListComprehension(list *ast.ListComprehension) object.Object {
	closeEnv := object.NewEnclosedEnvironment(ev.Env)

	res := &object.ListComprehension{
		Expression: list.Expression,
		Conditions: list.Conditions,
		IterateFunc: func(callback object.IterateCallback) {
			loopCallback := func(env *object.Environment) object.Object {
				val := ev.Eval(list.Expression, env)
				callback(val)
				return NULL
			}
			ev.evalForEach(list.Conditions, []ast.Expression{}, loopCallback, closeEnv)
		},
	}

	return res
}

func (ev *Evaluator) evalIntInterval(interval *ast.IntInterval) object.Object {
	lowerObj := ev.Eval(interval.Lower)
	upperObj := ev.Eval(interval.Upper)

	lower, ok1 := lowerObj.(object.Realness)
	upper, ok2 := upperObj.(object.Realness)
	if !ok1 {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm chặn dưới", lowerObj.Type())
		return ev.runtimeError(errMsg, interval.Lower)
	}
	if !ok2 {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm chặn trên", upperObj.Type())
		return ev.runtimeError(errMsg, interval.Upper)
	}
	if upper.Less(lower) == TRUE {
		errMsg := fmt.Sprintf("Chặn dưới phải nhỏ hơn chặn trên (%s > %s)", lower.Display(), upper.Display())
		return ev.runtimeError(errMsg)
	}

	return &object.IntInterval{Lower: lower, Upper: upper}
}

func (ev *Evaluator) evalRealInterval(interval *ast.RealInterval) object.Object {
	lowerObj := ev.Eval(interval.Lower)
	upperObj := ev.Eval(interval.Upper)

	lower, ok1 := lowerObj.(object.Realness)
	upper, ok2 := upperObj.(object.Realness)
	if !ok1 {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm chặn dưới", lowerObj.Type())
		return ev.runtimeError(errMsg)
	}
	if !ok2 {
		errMsg := fmt.Sprintf("Không thể dùng '%s' làm chặn trên", lowerObj.Type())
		return ev.runtimeError(errMsg)
	}

	return &object.RealInterval{Lower: lower, Upper: upper}
}
