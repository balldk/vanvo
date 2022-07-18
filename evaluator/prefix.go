package evaluator

import (
	"vila/object"
	"vila/token"
)

func evalPrefixExpression(operator token.Token, right object.Object) object.Object {
	switch operator.Type {
	case token.BANG:
		return evalBangPrefix(right)
	case token.MINUS:
		return evalMinusPrefix(right)
	case token.PLUS:
		return evalPlusPrefix(right)
	default:
		return NULL
	}
}

func evalBangPrefix(right object.Object) object.Object {
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

func evalMinusPrefix(right object.Object) object.Object {
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

func evalPlusPrefix(right object.Object) object.Object {
	return right
}
