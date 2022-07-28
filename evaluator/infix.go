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
	if left, isImply := left.(*object.Imply); isImply {
		return left
	}
	if right, isImply := right.(*object.Imply); isImply {
		return right
	}

	switch operator.Type {
	case token.Plus:
		return ev.evalAddition(operator, left, right)

	case token.Minus:
		return ev.evalSubtraction(operator, left, right)

	case token.Asterisk:
		return ev.evalMultiplication(operator, left, right)

	case token.Slash:
		return ev.evalDivision(operator, left, right)

	case token.Equal:
		return ev.evalEquality(operator, left, right)

	case token.Less:
		return ev.evalLess(operator, left, right)

	case token.Greater:
		return ev.evalLess(operator, right, left)

	case token.LessEqual:
		if ev.evalLess(operator, left, right) == TRUE {
			return TRUE
		}
		return ev.evalEquality(operator, left, right)

	case token.GreaterEqual:
		if ev.evalLess(operator, right, left) == TRUE {
			return TRUE
		}
		return ev.evalEquality(operator, left, right)

	case token.Hat:
		return NULL
	}

	return NULL
}

func (ev *Evaluator) evalAddition(operator token.Token, left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể cộng `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.Additive); ok {
		value := left.Add(right)
		if value == object.CANT_OPERATE {
			return ev.runtimeError(errMsg)
		}
		return value
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalSubtraction(operator token.Token, left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể trừ `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.Subtractive); ok {
		value := left.Subtract(right)
		if value == object.CANT_OPERATE {
			return ev.runtimeError(errMsg)
		}
		return value
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalMultiplication(operator token.Token, left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể nhân `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.Multiplicative); ok {
		value := left.Multiply(right)
		if value == object.CANT_OPERATE {
			return ev.runtimeError(errMsg)
		}
		return value
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalDivision(operator token.Token, left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể chia `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.Division); ok {
		value := left.Divide(right)
		if value == object.CANT_OPERATE {
			return ev.runtimeError(errMsg)
		}
		return value
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalEquality(operator token.Token, left, right object.Object) *object.Boolean {
	errMsg := fmt.Sprintf("Không thể so sánh `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.Equal); ok {
		value := left.Equal(right)
		if value == object.INCOMPARABLE {
			ev.runtimeError(errMsg)
			return INCOMPARABLE
		}
		return value
	}

	ev.runtimeError(errMsg)
	return INCOMPARABLE
}

func (ev *Evaluator) evalLess(operator token.Token, left, right object.Object) *object.Boolean {
	errMsg := fmt.Sprintf("Không thể so sánh `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.StrictOrder); ok {
		value := left.Less(right)
		if value == object.INCOMPARABLE {
			ev.runtimeError(errMsg)
			return INCOMPARABLE
		}
		return value
	}

	ev.runtimeError(errMsg)
	return INCOMPARABLE
}
