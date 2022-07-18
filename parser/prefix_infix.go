package parser

import (
	"vila/ast"
	"vila/token"
)

var precedences = map[token.TokenType]int{
	token.EQ:         EQUAL,
	token.NEQ:        EQUAL,
	token.LESS:       COMPARE,
	token.GREATER:    COMPARE,
	token.LESS_EQ:    COMPARE,
	token.GREATER_EQ: COMPARE,
	token.PLUS:       SUM,
	token.MINUS:      SUM,
	token.ASTERISK:   PRODUCT,
	token.DOT:        PRODUCT,
	token.SLASH:      PRODUCT,
	token.PERCENT:    PRODUCT,
	token.HAT:        EXP,
	token.LPAREN:     CALL,
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

	return expr
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.advanceToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}
