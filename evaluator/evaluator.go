package evaluator

import (
	"fmt"
	"vila/ast"
	"vila/errorhandler"
	"vila/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func New(errors *errorhandler.ErrorList) *Evaluator {
	ev := &Evaluator{Errors: errors}
	return ev
}

type Evaluator struct {
	Errors *errorhandler.ErrorList
	Node   ast.Node
}

func (ev *Evaluator) Eval(node ast.Node) object.Object {
	newev := New(ev.Errors)
	newev.Node = node
	return newev.evalNode()
}

func (ev *Evaluator) evalNode() object.Object {
	node := ev.Node

	if ev.Errors.NotEmpty() {
		return NULL
	}

	switch node := node.(type) {

	case *ast.Program:
		return ev.evalProgram(node)

	case *ast.ExpressionStatement:
		return ev.Eval(node.Expression)

	case *ast.PrefixExpression:
		right := ev.Eval(node.Right)
		return ev.evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := ev.Eval(node.Left)
		right := ev.Eval(node.Right)
		return ev.evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return ev.evalBlockStatement(node.Statements)

	case *ast.IfExpression:
		return ev.evalIfExpression(node)

	case *ast.ImplyStatement:
		val := ev.Eval(node.Value)
		return &object.Imply{Value: val}

	case *ast.Int:
		return &object.Int{Value: node.Value}

	case *ast.Real:
		return &object.Real{Value: node.Value}

	case *ast.Boolean:
		return boolRef(node.Value)

	}

	return NULL
}

func (ev *Evaluator) evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = ev.Eval(statement)

		if returnValue, ok := result.(*object.Imply); ok {
			return returnValue.Value
		}
	}
	return result
}

func (ev *Evaluator) evalBlockStatement(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = ev.Eval(statement)

		if result.Type() == object.IMPLY_OBJ {
			return result
		}
	}

	return result
}

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

func (ev *Evaluator) isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Null:
		return false
	case *object.Boolean:
		return obj.Value
	case *object.Int:
		if obj.Value == 0 {
			return false
		}
		return true
	case *object.Real:
		if obj.Value == 0 {
			return false
		}
		return true
	case *object.Quotient:
		if obj.Numer.Value == 0 {
			return false
		}
		return true
	default:
		errMsg := fmt.Sprintf("Không thể đặt `%s` làm điều kiện", obj.Type())
		ev.runtimeError(errMsg)
		return false
	}
}

func boolRef(val bool) *object.Boolean {
	if val {
		return TRUE
	}
	return FALSE
}

func (ev *Evaluator) runtimeError(msg string) object.Object {
	ev.Errors.AddRuntimeError(msg, ev.Node)
	return NULL
}
