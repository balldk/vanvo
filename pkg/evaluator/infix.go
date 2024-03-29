package evaluator

import (
	"fmt"
	"vanvo/pkg/object"
	"vanvo/pkg/token"
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

	case token.Percent:
		return ev.evalModulo(left, right)

	case token.Dot:
		return ev.evalDotProduct(left, right)

	case token.Hat:
		return ev.evalExponent(left, right)

	case token.Equal:
		return ev.evalEquality(left, right)

	case token.NotEqual:
		return ev.evalEquality(left, right).Not()

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

	case token.And:
		return ev.evalAnd(left, right)

	case token.Or:
		return ev.evalOr(left, right)

	case token.Belong:
		return ev.evalBelong(left, right)
	}

	return NULL
}

func (ev *Evaluator) evalAddition(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể cộng '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.Additive); ok {
		value := left.Add(right)
		return ev.someObject(value, errMsg)
	}
	if left, ok := left.(object.Set); ok {
		if right, ok := right.(object.Set); ok {
			return ev.evalUnion(left, right)
		}
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalSubtraction(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể trừ '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.Subtractive); ok {
		value := left.Subtract(right)
		return ev.someObject(value, errMsg)
	}
	if left, ok := left.(object.Set); ok {
		if right, ok := right.(object.Set); ok {
			return ev.evalSetDiff(left, right)
		}
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalMultiplication(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể nhân '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.Multiplicative); ok {
		value := left.Multiply(right)
		return ev.someObject(value, errMsg)
	}
	if left, ok := left.(object.Set); ok {
		if right, ok := right.(object.Set); ok {
			return ev.evalSetProduct(left, right)
		}
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalDivision(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể chia '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.Division); ok {
		value := left.Divide(right)
		return ev.someObject(value, errMsg)
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalModulo(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể chia lấy dư '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.Modulo); ok {
		value := left.Mod(right)
		return ev.someObject(value, errMsg)
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalDotProduct(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể . '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.DotProduct); ok {
		value := left.Dot(right)
		return ev.someObject(value, errMsg)
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalExponent(left, right object.Object) object.Object {
	errMsg := fmt.Sprintf("Không thể mũ '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.Exponential); ok {
		value := left.Power(right)
		return ev.someObject(value, errMsg)
	}
	if left, ok := left.(object.Set); ok {
		if right, ok := right.(*object.Int); ok {
			return ev.evalSetExponent(left, right)
		}
	}

	return ev.runtimeError(errMsg)
}

func (ev *Evaluator) evalUnion(left, right object.Set) object.Set {
	if left, isInterval := left.(*object.RealInterval); isInterval {
		if right, isInterval := right.(*object.RealInterval); isInterval {
			if right.Lower.Less(left.Lower) == TRUE {
				left, right = right, left
			}
			if right.Less(left).Value {
				return left
			}
			if !left.Upper.Less(right.Lower).Value {
				return &object.RealInterval{Lower: left.Lower, Upper: right.Upper}
			}
		}
	}

	if left, isList := left.(*object.List); isList {
		if right, isList := right.(*object.List); isList {
			return &object.List{Data: append(left.Data, right.Data...)}
		}
	}

	return &object.UnionSet{Left: left, Right: right}
}

func (ev *Evaluator) evalSetDiff(left, right object.Set) object.Set {
	return &object.DiffSet{Left: left, Right: right}
}

func (ev *Evaluator) evalSetProduct(left, right object.Set) object.Set {

	if left, isProd := left.(*object.ProductSet); isProd {
		return &object.ProductSet{Sets: append(left.Sets, right)}
	}
	return &object.ProductSet{Sets: []object.Set{left, right}}
}

func (ev *Evaluator) evalSetExponent(left object.Set, right *object.Int) object.Set {
	sets := []object.Set{}

	for i := 0; i < int(right.Value.Int64()); i++ {
		sets = append(sets, left)
	}
	return &object.ProductSet{Sets: sets}
}

func (ev *Evaluator) evalEquality(left, right object.Object) *object.Boolean {
	errMsg := fmt.Sprintf("Không thể so sánh '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.Equal); ok {
		value := left.Equal(right)
		value, isBool := ev.someObject(value, errMsg).(*object.Boolean)
		if !isBool {
			return FALSE
		}
		return value
	}

	ev.runtimeError(errMsg)
	return INCOMPARABLE
}

func (ev *Evaluator) evalLess(left, right object.Object) *object.Boolean {
	errMsg := fmt.Sprintf("Không thể so sánh '%v' với '%v'", left.Type(), right.Type())

	if left, ok := left.(object.StrictOrder); ok {
		value := left.Less(right)
		value, isBool := ev.someObject(value, errMsg).(*object.Boolean)
		if !isBool {
			return FALSE
		}
		return value
	}

	ev.runtimeError(errMsg)
	return INCOMPARABLE
}

func (ev *Evaluator) evalAnd(left, right object.Object) object.Object {
	if !ev.isTruthy(left) {
		return left
	}
	return right
}

func (ev *Evaluator) evalOr(left, right object.Object) object.Object {
	if ev.isTruthy(left) {
		return left
	}
	return right
}

func (ev *Evaluator) evalBelong(left, right object.Object) *object.Boolean {
	errMsg := fmt.Sprintf("Vế phải của mệnh đề 'thuộc' phải là một '%s' thay vì '%s'",
		object.SetObj, right.Type())

	if right, ok := right.(object.Set); ok {
		value := right.Contain(left)
		value, isBool := ev.someObject(value, errMsg).(*object.Boolean)
		if !isBool {
			return FALSE
		}
		return value
	}

	ev.runtimeError(errMsg)
	return INCOMPARABLE
}

func (ev *Evaluator) someObject(obj object.Object, errMsg string) object.Object {
	if obj == object.INCOMPARABLE || obj == object.CANT_OPERATE {
		return ev.runtimeError(errMsg)
	}
	if obj == object.ZERO_DIVISION {
		return ev.runtimeError("Không thể chia cho 0")
	}

	return obj
}
