package evaluator

import (
	"fmt"
	"vila/object"
	"vila/token"
)

func (ev *Evaluator) evalInfixExpression(
	operator token.Token,
	left, right object.Object,
) object.Object {

	switch operator.Type {
	case token.Plus:
		return ev.evalAddition(operator, left, right)
	case token.Minus:
		return NULL
	case token.Asterisk:
		return NULL
	case token.Slash:
		return NULL
	case token.Hat:
		return NULL
	}
	return NULL
}

func (ev *Evaluator) evalAddition(operator token.Token, left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể cộng `%v` với `%v`", left.Type(), right.Type())

	if left, ok1 := left.(object.Additive); ok1 {
		value := left.Add(right)
		if value == nil {
			ev.runtimeError(errMsg, operator)
		}

		return value
	}

	return ev.runtimeError(errMsg, operator)
}
