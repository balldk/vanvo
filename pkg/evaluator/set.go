package evaluator

import (
	"fmt"
	"vila/pkg/ast"
	"vila/pkg/object"
)

func (ev *Evaluator) evalIntInterval(interval *ast.IntInterval) object.Object {
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

	return &object.IntInterval{Lower: lower.ToReal(), Upper: upper.ToReal()}
}
