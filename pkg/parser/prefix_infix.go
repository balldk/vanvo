package parser

import (
	"vila/pkg/ast"
	"vila/pkg/token"
)

var precedences = map[token.TokenType]int{
	token.If:           IF,
	token.Equal:        EQUAL,
	token.NotEqual:     EQUAL,
	token.Less:         COMPARE,
	token.Greater:      COMPARE,
	token.LessEqual:    COMPARE,
	token.GreaterEqual: COMPARE,
	token.Plus:         SUM,
	token.Minus:        SUM,
	token.Asterisk:     PRODUCT,
	token.Dot:          PRODUCT,
	token.Slash:        PRODUCT,
	token.Percent:      PRODUCT,
	token.Hat:          EXP,
	token.LParen:       CALL,
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
		p.syntaxErrorImportant("Thiếu vế phải của " + string(expr.Operator.Literal))
	}

	return expr
}
