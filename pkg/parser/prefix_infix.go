package parser

import (
	"vanvo/pkg/ast"
	"vanvo/pkg/token"
)

var precedences = map[token.TokenType]int{
	token.If:           IF,
	token.And:          CONJUNC,
	token.Or:           CONJUNC,
	token.Belong:       BELONG,
	token.Equal:        EQUAL,
	token.NotEqual:     EQUAL,
	token.Less:         COMPARE,
	token.Greater:      COMPARE,
	token.LessEqual:    COMPARE,
	token.GreaterEqual: COMPARE,
	token.Plus:         SUM,
	token.Minus:        SUM,
	token.Asterisk:     PRODUCT,
	token.Slash:        PRODUCT,
	token.Percent:      PRODUCT,
	token.Hat:          EXP,
	token.LParen:       CALL,
	token.LBracket:     CALL,
	token.Dot:          Compose,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) peekPrecedence() int {
	if pre, ok := precedences[p.peekToken.Type]; ok {
		return pre
	}

	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if prec, ok := precedences[p.curToken.Type]; ok {
		return prec
	}

	return LOWEST
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Operator: p.curToken,
	}

	p.advanceToken()
	expr.Right = p.parseExpression(PREFIX)

	if expr.Right == nil {
		p.syntaxError("Tiền tố không tồn tại")
	}

	return expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Operator: p.curToken,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.advanceToken()
	expr.Right = p.parseExpression(precedence)

	if expr.Right == nil {
		p.syntaxError("Thiếu vế phải của " + string(expr.Operator.Literal))
	}

	// convert expression like a < b < c to (a < b) và (b < c)
	if left, isInfix := left.(*ast.InfixExpression); isInfix {
		cond1 := left.Operator.Type == token.Less || left.Operator.Type == token.LessEqual
		cond2 := expr.Operator.Type == token.Less || expr.Operator.Type == token.LessEqual

		cond3 := left.Operator.Type == token.Greater || left.Operator.Type == token.GreaterEqual
		cond4 := expr.Operator.Type == token.Greater || expr.Operator.Type == token.GreaterEqual

		cond5 := left.Operator.Type == token.Equal && expr.Operator.Type == token.Equal

		if (cond1 && cond2) || (cond3 && cond4) || cond5 {
			expr = &ast.InfixExpression{
				Operator: token.Token{Type: token.And},
				Left:     left,
				Right: &ast.InfixExpression{
					Operator: expr.Operator,
					Left:     left.Right,
					Right:    expr.Right,
				},
			}
		}
	}

	return expr
}
