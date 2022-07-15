package parser

import (
	"strconv"
	"vila/ast"
	"vila/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX
	CALL
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.Errors.AddSyntaxError("tiền tố không tồn tại", &p.curToken)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseInt() ast.Expression {
	i := &ast.Int{Token: p.curToken}

	value, err := strconv.ParseInt(string(p.curToken.Literal), 0, 64)
	if err != nil {
		p.Errors.AddSyntaxError("Không thể parse số nguyên này", &p.curToken)
	}

	i.Value = value

	return i
}

func (p *Parser) parseReal() ast.Expression {
	re := &ast.Real{Token: p.curToken}

	value, err := strconv.ParseFloat(string(p.curToken.Literal), 64)
	if err != nil {
		p.Errors.AddSyntaxError("Không thể parse số thực này", &p.curToken)
	}

	re.Value = value

	return re
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}
