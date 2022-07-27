package parser

import (
	"vila/ast"
	"vila/token"
)

func (p *Parser) parseStatement() ast.Statement {
	p.updateIdentLevel()

	var stmt ast.Statement

	if p.curTokenIs(token.Ident) && p.peekTokenIs(token.Assign) {
		stmt = p.parseAssignStatement()
	}

	switch p.curToken.Type {
	case token.Let:
		stmt = p.parseLetStatement()
	case token.If:
		stmt = p.parseIfStatement()
		p.updateIdentLevel()
		return stmt
	case token.Imply:
		stmt = p.parseImplyStatement()
	default:
		stmt = p.parseExpressionStatement()
	}

	p.checkEndStatement()
	p.updateIdentLevel()

	return stmt
}

func (p *Parser) checkEndStatement() {
	p.advanceToken()
	if !p.curIsStatementSeperator() && !p.curTokenIs(token.RParen) {
		p.invalidSyntax()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{}
	stmt.Expression = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// Tên biến
	if !p.expectPeek(token.Ident) {
		p.Errors.AddParserError("Sau `cho` phải là một tên định danh", p.curToken)
	}
	stmt.Ident = &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}

	// If this is function
	// if p.peekTokenIs(token.LParen) {
	// 	return p.parseFunction()
	// }

	hasAssign := false

	// If has assign
	if p.peekTokenIs(token.Assign) {
		hasAssign = true
		p.advanceToken()
		p.advanceToken()

		stmt.Value = p.parseExpression(LOWEST)
	}

	// If has belong clause
	if p.peekTokenIs(token.Belong) {
		p.advanceToken()
		p.advanceToken()

		stmt.SetType = p.parseExpression(LOWEST)

	} else if !hasAssign {
		p.Errors.AddParserError("Khai báo biến phải có giá trị khởi tạo hoặc miền xác định", p.curToken)
	}

	return stmt
}

func (p *Parser) parseImplyStatement() *ast.ImplyStatement {
	stmt := &ast.ImplyStatement{Token: p.curToken}
	p.advanceToken()

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}
	p.advanceToken()

	stmt.Condition = p.parseExpression(LOWEST)
	stmt.Consequence = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	if !p.expectPeek(token.Colon) {
		return nil
	}

	block := &ast.BlockStatement{Colon: p.curToken}
	block.Statements = []ast.Statement{}
	p.advanceToken()

	p.identLevel++
	curLevel := p.identLevel

	for p.identLevel == curLevel && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		block.Statements = append(block.Statements, stmt)
		p.updateIdentLevel()
	}

	return block
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{Token: p.curToken}
	stmt.Ident = &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}

	p.advanceToken()
	p.advanceToken()

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}
