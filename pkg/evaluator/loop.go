package evaluator

import (
	"fmt"
	"vila/pkg/ast"
	"vila/pkg/object"
	"vila/pkg/token"
)

func (ev *Evaluator) evalForEachStatement(
	stmt *ast.ForEachStatement,
	env *object.Environment,
	constraints []ast.Expression,
) object.Object {

	// if no condition left
	if len(stmt.Conditions) == 0 {
		for _, cons := range constraints {
			check := ev.Eval(cons, env)
			if !ev.isTruthy(check) {
				return NULL
			}
		}
		return ev.Eval(stmt.Body, env)
	}

	// if current condition is 'belong' clause
	if condition, ok := stmt.Conditions[0].(*ast.InfixExpression); ok {
		if condition.Operator.Type == token.Belong {

			right := ev.Eval(condition.Right)
			loopSet, isCountable := right.(object.CountableSet)
			if !isCountable || !loopSet.IsCountable() {
				errMsg := fmt.Sprintf("Vế phải của mệnh đề 'thuộc' phải là một 'Tập đếm được' thay vì '%s'", right.Type())
				return ev.runtimeError(errMsg, condition.Right)
			}

			ident, isIdent := condition.Left.(*ast.Identifier)
			if !isIdent {
				errMsg := "Vế trái của mệnh đề 'thuộc' phải là một tên định danh"
				return ev.runtimeError(errMsg, condition.Left)
			}

			loopSet.Iterate(func(element object.Object) {
				env.SetInScope(ident.Value, element)

				fullConditions := stmt.Conditions
				stmt.Conditions = stmt.Conditions[1:]

				closeEnv := object.NewEnclosedEnvironment(env)
				ev.evalForEachStatement(stmt, closeEnv, constraints)

				stmt.Conditions = fullConditions
			})

			return NULL
		}
	}

	// constraints
	constraints = append(constraints, stmt.Conditions[0])

	stmt.Conditions = stmt.Conditions[1:]

	closeEnv := object.NewEnclosedEnvironment(env)
	ev.evalForEachStatement(stmt, closeEnv, constraints)

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
