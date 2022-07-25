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
		return ev.evalSubtraction(operator, left, right)
	case token.Asterisk:
		return ev.evalMultiplication(operator, left, right)
	case token.Slash:
		return ev.evalDivision(operator, left, right)
	case token.Hat:
		return NULL
	}
	return NULL
}

func (ev *Evaluator) evalAddition(operator token.Token, left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể cộng `%v` với `%v`", left.Type(), right.Type())

	if left, ok := left.(object.Additive); ok {
		value := left.Add(right)
		if value == nil {
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
		if value == nil {
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
		if value == nil {
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
		if value == nil {
			return ev.runtimeError(errMsg)
		}
		return value
	}

	return ev.runtimeError(errMsg)
}

func SomeObject(value object.Object) object.Object {
	if value == nil {
		return NULL
	}
	return value
}
