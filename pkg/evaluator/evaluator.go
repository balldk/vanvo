package evaluator

import (
	"fmt"
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

	case *ast.ImplyStatement:
		val := ev.Eval(node.Value)
		return &object.Imply{Value: val}

	case *ast.AssignStatement:
		val := ev.Eval(node.Value)
		obj := ev.Env.Set(node.Ident.Value, val)
		if obj == nil {
			errMsg := fmt.Sprintf("`%s` chưa được khai báo", node.Ident.Value)
			return ev.runtimeError(errMsg)
		}

	case *ast.VarDeclareStatement:
		val := ev.Eval(node.Value)
		if _, ok := ev.Env.GetInScope(node.Ident.Value); ok {
			errMsg := fmt.Sprintf("'%s' đã được khởi tạo", node.Ident.Value)
			ev.runtimeError(errMsg)
		}
		ev.Env.SetInScope(node.Ident.Value, val)

	case *ast.FunctionDeclareStatement:
		params := node.Params
		body := node.Body
		fn := &object.Function{Ident: node.Ident, Params: params, Body: body}
		if _, ok := ev.Env.GetInScope(node.Ident.Value); ok {
			errMsg := fmt.Sprintf("'%s' đã được khởi tạo", node.Ident.Value)
			ev.runtimeError(errMsg)
		}
		ev.Env.SetInScope(node.Ident.Value, fn)

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

func (ev *Evaluator) evalIfStatement(ie *ast.IfStatement) object.Object {
	for _, branch := range ie.Branches {

		condition := ev.Eval(branch.Condition)
		if ev.isTruthy(condition) {
			return ev.Eval(branch.Consequence)
		}
	}

	return NULL
}

func (ev *Evaluator) evalCallExpression(call *ast.CallExpression) object.Object {
	fn := ev.Eval(call.Function)
	args := ev.evalExpressions(call.Arguments)

	switch fn := fn.(type) {
	case *object.Function:
		res := ev.applyFunction(fn, args)
		args = []object.Object{res}

		for fn.LeftCompose != nil {
			fn = fn.LeftCompose
			res = ev.applyFunction(fn, args)
			args = []object.Object{res}
		}

		return res

	default:
		if len(call.Arguments) == 1 {
			right := ev.Eval(call.Arguments[0])
			return ev.evalMultiplication(fn, right)
		}

		return ev.runtimeError("Biểu thức không hợp lệ")
	}
}

func (ev *Evaluator) applyFunction(fn *object.Function, args []object.Object) object.Object {
	env := object.NewEnclosedEnvironment(ev.Env)

	if fn.Builtin != nil {
		return fn.Builtin(args...)
	}

	if len(args) != len(fn.Params) {
		errMsg := fmt.Sprintf(
			"'%s' cần %d tham số thay vì %d",
			fn.Ident.Value, len(fn.Params), len(args))

		return ev.runtimeError(errMsg)
	}

	for index, param := range fn.Params {
		env.SetInScope(param.Value, args[index])
	}

	val := ev.Eval(fn.Body, env)
	return ev.unwrapImply(val)
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
		fmt.Print(evaluated.Display(), " ")
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

func (ev *Evaluator) unwrapImply(obj object.Object) object.Object {
	if imply, ok := obj.(*object.Imply); ok {
		return ev.unwrapImply(imply.Value)
	}

	return obj
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
