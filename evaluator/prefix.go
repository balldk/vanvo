package evaluator

import (
	"vila/object"
	"vila/token"
)

func (ev *Evaluator) evalPrefixExpression(operator token.Token, right object.Object) object.Object {
	switch operator.Type {
	case token.BANG:
		return ev.evalBangPrefix(right)
	case token.MINUS:
		return ev.evalMinusPrefix(right)
	case token.PLUS:
		return ev.evalPlusPrefix(right)
	default:
		return NULL
	}
}

func (ev *Evaluator) evalBangPrefix(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func (ev *Evaluator) evalMinusPrefix(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Int:
		right.Value = -right.Value
		return right
	case *object.Real:
		right.Value = -right.Value
		return right
	default:
		return NULL
	}
}

func (ev *Evaluator) evalPlusPrefix(right object.Object) object.Object {
	return right
}
