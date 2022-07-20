package evaluator

import (
	"vila/ast"
	"vila/errorhandler"
	"vila/object"
	"vila/token"
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
}

func (ev *Evaluator) Eval(node ast.Node) object.Object {
	if ev.Errors.NotEmpty() {
		return NULL
	}

	switch node := node.(type) {

	case *ast.Program:
		return ev.evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return ev.Eval(node.Expression)

	case *ast.PrefixExpression:
		right := ev.Eval(node.Right)
		return ev.evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := ev.Eval(node.Left)
		right := ev.Eval(node.Right)
		return ev.evalInfixExpression(node.Operator, left, right)

	case *ast.Int:
		return &object.Int{Value: node.Value}

	case *ast.Real:
		return &object.Real{Value: node.Value}

	case *ast.Boolean:
		return boolRef(node.Value)

	}

	return NULL
}

func (ev *Evaluator) evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = ev.Eval(statement)
	}

	return result
}

func boolRef(val bool) *object.Boolean {
	if val {
		return TRUE
	}
	return FALSE
}

func (ev *Evaluator) runtimeError(msg string, tok token.Token) object.Object {
	ev.Errors.AddRuntimeError(msg, tok)
	return NULL
}
