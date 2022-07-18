package evaluator

import (
	"vila/object"
	"vila/token"
)

func evalInfixExpression(
	operator token.Token,
	left, right object.Object,
) object.Object {

	switch operator.Type {
	case token.PLUS:
		return evalAddition(left, right)
	}
	return NULL
}

func evalAddition(left, right object.Object) object.Object {
	if left, ok1 := left.(object.Additive); ok1 {
		if right, ok2 := right.(object.Additive); ok2 {
			left.Add(right)
		}
	}
	return NULL
}
