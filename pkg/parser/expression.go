package parser

import (
	"strconv"
	"vila/pkg/ast"
	"vila/pkg/token"
)

const (
	_ int = iota
	LOWEST
	IF
	EQUAL   // ==
	COMPARE // > or <
	SUM     // +
	PRODUCT // *
	EXP     // ^
	PREFIX
	CALL
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.invalidSyntax()
		return nil
	}
	leftExp := prefix()

	for !p.peekIsStatementSeperator() && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			p.syntaxError("toán tử trung tố không tồn tại")
			return leftExp
		}

		p.advanceToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}
}

func (p *Parser) parseInt() ast.Expression {
	i := &ast.Int{Token: p.curToken}

	value, err := strconv.ParseInt(string(p.curToken.Literal), 0, 64)
	if err != nil {
		p.syntaxError("Không thể parse số nguyên này")
	}

	i.Value = value

	if p.peekTokenIs(token.Ident) {
		p.insertPeekToken(token.Token{
			Type:    token.Asterisk,
			Literal: []rune("*"),
			Line:    p.peekToken.Line,
			Column:  p.peekToken.Column,
		})
	}

	return i
}

func (p *Parser) parseReal() ast.Expression {
	re := &ast.Real{Token: p.curToken}

	value, err := strconv.ParseFloat(string(p.curToken.Literal), 64)
	if err != nil {
		p.syntaxError("Không thể parse số thực này")
	}

	re.Value = value

	if p.peekTokenIs(token.Ident) {
		p.insertPeekToken(token.Token{
			Type:    token.Asterisk,
			Literal: []rune("*"),
			Line:    p.peekToken.Line,
			Column:  p.peekToken.Column,
		})
	}

	return re
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.True)}
}

func (p *Parser) parseInterval() ast.Expression {
	p.advanceToken()
	leftBracket := p.curToken

	lower := p.parseExpression(LOWEST)

	p.advanceToken()
	if p.curTokenIs(token.Comma) { // real interval
		p.advanceToken()

		seg := &ast.RealInterval{
			LeftBracket: leftBracket,
			Lower:       lower,
			Upper:       p.parseExpression(LOWEST),
		}
		if p.expectPeek(token.RBracket) {
			seg.RightBracket = p.curToken
			return seg
		}

	} else if p.curTokenIs(token.DotDot) { // int interval
		p.advanceToken()

		seg := &ast.IntInterval{
			LeftBracket: leftBracket,
			Lower:       lower,
			Upper:       p.parseExpression(LOWEST),
		}
		if p.expectPeek(token.RBracket) {
			seg.RightBracket = p.curToken
			return seg
		}
	} else {
		p.invalidSyntax()
	}

	return nil
}

func (p *Parser) parseIfExpression(left ast.Expression) ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken, Consequence: left}

	p.advanceToken()
	expression.Condition = p.parseExpression(LOWEST)

	if p.expectPeek(token.Else) {
		p.advanceToken()
		expression.Alternative = p.parseExpression(LOWEST)
	}

	return expression
}

func (p *Parser) parseGroupExpression() ast.Expression {
	block := &ast.GroupExpression{LeftParen: p.curToken}
	block.Statements = []ast.Statement{}

	if !p.expectCur(token.LParen) {
		return nil
	}

	if p.curTokenIs(token.Endline) {
		p.identLevel++
	}

	for !p.curTokenIs(token.RParen) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		block.Statements = append(block.Statements, stmt)
	}

	if !p.curTokenIs(token.RParen) {
		p.expectError(token.RParen)
		return nil
	}

	block.RightParen = p.curToken

	return block
}

func (p *Parser) parseCallExpression(fn ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Function: fn}
	exp.Arguments = p.parseCallArguments()
	exp.RightParen = p.curToken

	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RParen) {
		p.advanceToken()
		return args
	}

	p.advanceToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.Comma) {
		p.advanceToken()
		p.advanceToken()
		if p.curTokenIs(token.RParen) {
			return args
		}
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RParen) {
		return nil
	}

	return args
}