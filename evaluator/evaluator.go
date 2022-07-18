package evaluator

import (
	"vila/ast"
	"vila/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.Int:
		return &object.Int{Value: node.Value}

	case *ast.Real:
		return &object.Real{Value: node.Value}

	case *ast.Boolean:
		return boolRef(node.Value)

	}

	return NULL
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

func boolRef(val bool) *object.Boolean {
	if val {
		return TRUE
	}
	return FALSE
}
