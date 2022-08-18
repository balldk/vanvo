package evaluator

import (
	"vanvo/pkg/ast"
	"vanvo/pkg/object"
	"vanvo/pkg/token"
)

func (ev *Evaluator) evalForEachStatement(stmt *ast.ForEachStatement) object.Object {
	closeEnv := object.NewEnclosedEnvironment(ev.Env)
	callback := func(loopEnv *object.Environment) object.Object {
		return ev.Eval(stmt.Body, loopEnv)
	}

	return ev.evalForEach(stmt.Conditions, []ast.Expression{}, callback, closeEnv)
}

func (ev *Evaluator) evalForEach(
	rawConditions []ast.Expression,
	constraints []ast.Expression,
	callback func(*object.Environment) object.Object,
	env *object.Environment,
) object.Object {

	var result object.Object
	result = NULL

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
				errMsg := "Vế phải của mệnh đề 'thuộc' phải là một Tập đếm được"
				return ev.runtimeError(errMsg, condition.Right)
			}

			loopSet.Iterate(func(element object.Object) object.Object {
				if result.Type() == object.IMPLY_OBJ {
					return result
				}
				env.SetInScope(ident.Value, element)

				fullConditions := rawConditions
				rawConditions = rawConditions[1:]

				closeEnv := object.NewEnclosedEnvironment(env)
				result = ev.evalForEach(rawConditions, constraints, callback, closeEnv)

				rawConditions = fullConditions

				return result
			})

			return result
		}
	}

	// constraints
	constraints = append(constraints, rawConditions[0])

	rawConditions = rawConditions[1:]

	closeEnv := object.NewEnclosedEnvironment(env)
	return ev.evalForEach(rawConditions, constraints, callback, closeEnv)
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
