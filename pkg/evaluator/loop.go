package evaluator

import (
	"fmt"
	"vila/pkg/ast"
	"vila/pkg/object"
	"vila/pkg/token"
)

func (ev *Evaluator) evalForEachStatement(stmt *ast.ForEachStatement) {
	closeEnv := object.NewEnclosedEnvironment(ev.Env)
	callback := func(loopEnv *object.Environment) object.Object {
		return ev.Eval(stmt.Body, loopEnv)
	}

	ev.evalForEach(stmt.Conditions, []ast.Expression{}, callback, closeEnv)
}

func (ev *Evaluator) evalForEach(
	rawConditions []ast.Expression,
	constraints []ast.Expression,
	callback func(*object.Environment) object.Object,
	env *object.Environment,
) object.Object {

	// if no condition left
	if len(rawConditions) == 0 {
		for _, cons := range constraints {
			check := ev.Eval(cons, env)
			if !ev.isTruthy(check) {
				return NULL
			}
		}
		return callback(env)
	}

	// if current condition is 'belong' clause
	if condition, ok := rawConditions[0].(*ast.InfixExpression); ok &&
		condition.Operator.Type == token.Belong {

		if ident, isIdent := condition.Left.(*ast.Identifier); isIdent {

			right := ev.Eval(condition.Right)
			loopSet, isCountable := right.(object.CountableSet)
			if !isCountable || !loopSet.IsCountable() {
				errMsg := fmt.Sprintf("Vế phải của mệnh đề 'thuộc' phải là một 'Tập đếm được' thay vì '%s'", right.Type())
				return ev.runtimeError(errMsg, condition.Right)
			}

			loopSet.Iterate(func(element object.Object) {
				env.SetInScope(ident.Value, element)

				fullConditions := rawConditions
				rawConditions = rawConditions[1:]

				closeEnv := object.NewEnclosedEnvironment(env)
				ev.evalForEach(rawConditions, constraints, callback, closeEnv)

				rawConditions = fullConditions
			})

			return NULL
		}
	}

	// constraints
	constraints = append(constraints, rawConditions[0])

	rawConditions = rawConditions[1:]

	closeEnv := object.NewEnclosedEnvironment(env)
	ev.evalForEach(rawConditions, constraints, callback, closeEnv)

	return NULL
}

func (ev *Evaluator) evalForStatement(stmt *ast.ForStatement) object.Object {
	for {
		for _, cond := range stmt.Conditions {
			check := ev.Eval(cond)
			if !ev.isTruthy(check) {
				return NULL
			}
		}
		ev.Eval(stmt.Body)
	}
}
