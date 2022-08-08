package evaluator

import (
	"math/big"
	"vila/pkg/object"
	"vila/pkg/token"
)

func (ev *Evaluator) evalPrefixExpression(operator token.Token, right object.Object) object.Object {
	if right, isImply := right.(*object.Imply); isImply {
		return right
	}

	switch operator.Type {
	case token.Bang:
		return ev.evalBangPrefix(right)
	case token.Minus:
		return ev.evalMinusPrefix(right)
	case token.Plus:
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
		return object.NewInt(new(big.Int).Neg(right.Value))
	case *object.Real:
		return object.NewReal(new(big.Float).Neg(right.Value))
	case *object.Quotient:
		numer := new(big.Int).Neg(right.Value.Num())
		return object.NewQuotient(numer, right.Value.Denom())
	case *object.Complex:
		return object.NewReal(object.RealZero).Subtract(right)
	default:
		return NULL
	}
}

func (ev *Evaluator) evalPlusPrefix(right object.Object) object.Object {
	return right
}
