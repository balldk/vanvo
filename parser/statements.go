package parser

import (
	"vila/ast"
	"vila/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.IMPLY:
		return p.parseImplyStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekIsStatementSeperator() {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// Tên biến
	if !p.expectPeek(token.IDENT) {
		p.Errors.AddSyntaxError("Sau 'cho' phải là một tên định danh", &p.curToken)
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	hasAssign := false

	// Nếu có lệnh gán
	if p.peekTokenIs(token.ASSIGN) {
		hasAssign = true
		p.nextToken()
		p.nextToken()

		stmt.Value = p.parseExpression(LOWEST)

	}

	// Nếu có mệnh đề 'thuộc'
	if p.peekTokenIs(token.BELONG) {
		p.nextToken()
		p.nextToken()

		stmt.SetType = p.parseExpression(LOWEST)

	} else if !hasAssign {
		p.Errors.AddSyntaxError("Khai báo biến phải có giá trị khởi tạo hoặc miền xác định", &p.curToken)
	}

	return stmt
}

func (p *Parser) parseImplyStatement() *ast.ImplyStatement {
	stmt := &ast.ImplyStatement{Token: p.curToken}
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if !p.curIsStatementSeperator() {
		p.nextToken()
	}

	return stmt
}
