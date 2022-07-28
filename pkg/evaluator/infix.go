package evaluator

import (
	"fmt"
	"vila/pkg/object"
	"vila/pkg/token"
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
		return ev.evalAddition(left, right)

	case token.Minus:
		return ev.evalSubtraction(left, right)

	case token.Asterisk:
		return ev.evalMultiplication(left, right)

	case token.Slash:
		return ev.evalDivision(left, right)

	case token.Equal:
		return ev.evalEquality(left, right)

	case token.Less:
		return ev.evalLess(left, right)

	case token.Greater:
		return ev.evalLess(right, left)

	case token.LessEqual:
		if ev.evalLess(left, right) == TRUE {
			return TRUE
		}
		return ev.evalEquality(left, right)

	case token.GreaterEqual:
		if ev.evalLess(right, left) == TRUE {
			return TRUE
		}
		return ev.evalEquality(left, right)

	case token.Hat:
		return ev.evalExponent(left, right)
	}

	return NULL
}

func (ev *Evaluator) evalAddition(left, right object.Object) object.Object {
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

func (ev *Evaluator) evalSubtraction(left, right object.Object) object.Object {
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

func (ev *Evaluator) evalMultiplication(left, right object.Object) object.Object {
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

func (ev *Evaluator) evalDivision(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể chia `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.Division); ok {
		value := left.Divide(right)
		if value == object.CANT_OPERATE {
			return ev.runtimeError(errMsg)

		} else if value == object.ZERO_DIVISION {
			return ev.runtimeError("Không thể chia cho 0")
		}
		return value
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalExponent(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể mũ `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.Exponential); ok {
		value := left.Power(right)
		if value == object.CANT_OPERATE {
			return ev.runtimeError(errMsg)
		}
		return value
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalEquality(left, right object.Object) *object.Boolean {
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

func (ev *Evaluator) evalLess(left, right object.Object) *object.Boolean {
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
