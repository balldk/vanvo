package evaluator

import (
	"fmt"
	"math/big"
	"vila/pkg/ast"
	"vila/pkg/errorhandler"
	"vila/pkg/lexer"
	"vila/pkg/object"
	"vila/pkg/parser"
)

var (
	NULL         = object.NULL
	TRUE         = object.TRUE
	FALSE        = object.FALSE
	INCOMPARABLE = object.INCOMPARABLE
	NO_PRINT     = &object.Null{}
)

func EvalFromInput(
	input string,
	path string,
	env *object.Environment,
) (object.Object, *errorhandler.ErrorList) {

	errors := errorhandler.NewErrorList(input, path)

	l := lexer.New(input, errors)
	p := parser.New(l, errors)
	ev := New(env, errors)

	program := p.ParseProgram()

	value := ev.Eval(program)

	return value, errors
}

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

	case *ast.GroupExpression:
		return ev.evalBlockStatement(node.Statements)

	case *ast.BlockStatement:
		return ev.evalBlockStatement(node.Statements)

	case *ast.IfExpression:
		return ev.evalIfExpression(node)

	case *ast.IfStatement:
		return ev.evalIfStatement(node)

	case *ast.ForStatement:
		return ev.evalForStatement(node)

	case *ast.ForEachStatement:
		return ev.evalForEachStatement(node)

	case *ast.ImplyStatement:
		val := ev.Eval(node.Value)
		return &object.Imply{Value: val}

	case *ast.AssignStatement:
		return ev.evalAssignStatement(node)

	case *ast.VarDeclareStatement:
		return ev.evalVarDeclare(node)

	case *ast.FunctionDeclareStatement:
		ev.evalFunctionDeclare(node)

	case *ast.OutputStatement:
		ev.evalOutputStatement(node)

	case *ast.CallExpression:
		return ev.evalCallExpression(node)

	case *ast.Identifier:
		return ev.evalIdentifier(node)

	case *ast.Int:
		return &object.Int{Value: node.Value}

	case *ast.Real:
		return &object.Real{Value: node.Value}

	case *ast.Boolean:
		return boolRef(node.Value)

	case *ast.String:
		return &object.String{Value: node.Value}

	case *ast.IndexExpression:
		return ev.evalIndex(node)

	case *ast.List:
		return ev.evalList(node)

	case *ast.ListComprehension:
		return ev.evalListComprehension(node)

	case *ast.IntInterval:
		return ev.evalIntInterval(node)

	case *ast.RealInterval:
		return ev.evalRealInterval(node)
	}

	return NO_PRINT
}

func (ev *Evaluator) evalProgram(program *ast.Program) object.Object {
	var result object.Object
	result = NO_PRINT

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
		errMsg := fmt.Sprintf("'%s' chưa được khởi tạo", node.Value)
		return ev.runtimeError(errMsg)
	}

	return val
}

func (ev *Evaluator) evalExpressions(exps []ast.Expression) []object.Object {
	var results []object.Object

	for _, e := range exps {
		result := ev.Eval(e)
		results = append(results, result)
	}
	return results
}

func (ev *Evaluator) evalOutputStatement(stmt *ast.OutputStatement) {
	for _, value := range stmt.Values {
		evaluated := ev.Eval(value)
		if str, isString := evaluated.(*object.String); isString {
			fmt.Print(str.Value, " ")
		} else {
			fmt.Print(evaluated.Display(), " ")
		}
	}
	fmt.Println()
}

func (ev *Evaluator) isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Null:
		return false
	case *object.Boolean:
		return obj.Value
	case *object.Int:
		if obj.Value.Cmp(object.IntZero) == 0 {
			return false
		}
		return true
	case *object.Real:
		if obj.Value.Cmp(big.NewFloat(0)) == 0 {
			return false
		}
		return true
	case *object.Quotient:
		if obj.Value.Num().Cmp(object.IntZero) == 0 {
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

func (ev *Evaluator) runtimeError(msg string, asts ...ast.Node) object.Object {
	node := ev.Node
	if len(asts) > 0 {
		node = asts[0]
	}
	ev.Errors.AddRuntimeError(msg, node)
	return NULL
}
