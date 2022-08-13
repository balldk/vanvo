package evaluator

import (
	"fmt"
	"vila/pkg/ast"
	"vila/pkg/object"
)

func (ev *Evaluator) evalIndex(exp *ast.IndexExpression) object.Object {
	set := ev.Eval(exp.Set)
	index := ev.Eval(exp.Index)

	if set, ok := set.(object.CountableSet); ok {
		if !set.IsCountable() {
			return ev.runtimeError("Chỉ 'Tập đếm được' mới có thể truy cập chỉ số")
		}

		if index, ok := index.(*object.Int); ok {
			val := set.At(int(index.Value.Int64()))

			if val == object.IndexError {
				return ev.runtimeError("Chỉ số vượt quá độ dài của tập hợp")
			}
			return val
		}
		errMsg := fmt.Sprintf("Chỉ số phải là một '%v' thay vì '%v'", object.IntObj, index.Type())
		return ev.runtimeError(errMsg)
	}
	return ev.runtimeError("Chỉ 'Tập đếm được' mới có thể truy cập chỉ số")
}

func (ev *Evaluator) evalList(list *ast.List) object.Object {
	exps := ev.evalExpressions(list.Data)
	return &object.List{Data: exps}
}

func (ev *Evaluator) evalListComprehension(node *ast.ListComprehension) object.Object {
	closeEnv := object.NewEnclosedEnvironment(ev.Env)

	list := &object.ListComprehension{
		Expression: node.Expression,
		Conditions: node.Conditions,
		Channel:    make(chan object.Object),
		Data:       []object.Object{},
	}
	go func() {
		defer close(list.Channel)
		callback := func(env *object.Environment) object.Object {
			val := ev.Eval(node.Expression, env)
			list.Data = append(list.Data, val)
			list.Channel <- val
			return val
		}
		ev.evalForEach(node.Conditions, []ast.Expression{}, callback, closeEnv)
	}()

	return list
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

	return &object.IntInterval{Lower: lower, Upper: upper, Step: object.NewInt(object.IntOne)}
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
