package evaluator

import (
	"vila/pkg/ast"
	"vila/pkg/object"
)

func (ev *Evaluator) evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := ev.Eval(ie.Condition)

	if ev.isTruthy(condition) {
		return ev.Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return ev.Eval(ie.Alternative)
	} else {
		return NULL
	}
}

func (ev *Evaluator) evalIfStatement(ie *ast.IfStatement) object.Object {
	for _, branch := range ie.Branches {

		condition := ev.Eval(branch.Condition)
		if ev.isTruthy(condition) {
			return ev.Eval(branch.Consequence)
		}
	}

	return NULL
}
