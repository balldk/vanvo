package evaluator

import (
	"fmt"
	"vila/ast"
	"vila/errorhandler"
	"vila/object"
)

var (
	NULL         = object.NULL
	TRUE         = object.TRUE
	FALSE        = object.FALSE
	INCOMPARABLE = object.INCOMPARABLE
)

func New(env *object.Environment, errors *errorhandler.ErrorList) *Evaluator {
	ev := &Evaluator{Errors: errors, Env: env}
	return ev
}

type Evaluator struct {
	Errors *errorhandler.ErrorList
	Node   ast.Node
	Env    *object.Environment
}

func (ev *Evaluator) Eval(node ast.Node, envs ...*object.Environment) object.Object {
	env := ev.Env
	if len(envs) > 0 {
		env = envs[0]
	}
	newev := New(env, ev.Errors)
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

	case *ast.LetStatement:
		val := ev.Eval(node.Value)
		ev.Env.Set(node.Ident.Value, val)

	case *ast.Identifier:
		return ev.evalIdentifier(node)

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
	env := object.NewEnclosedEnvironment(ev.Env)

	for _, statement := range stmts {
		result = ev.Eval(statement, env)

		if result.Type() == object.IMPLY_OBJ {
			return result
		}
	}

	return result
}

func (ev *Evaluator) evalIdentifier(node *ast.Identifier) object.Object {
	val, ok := ev.Env.Get(node.Value)
	if !ok {
		errMsg := fmt.Sprintf("Biến `%s` chưa được định nghĩa", node.Value)
		return ev.runtimeError(errMsg)
	}

	return val
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
