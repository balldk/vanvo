package parser

import (
	"vila/ast"
	"vila/token"
)

func (p *Parser) parseStatement() ast.Statement {
	if p.curTokenIs(token.Ident) && p.peekTokenIs(token.Assign) {
		return p.parseAssignStatement()
	}

	switch p.curToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Imply:
		return p.parseImplyStatement()
	default:
		return p.parseExpressionStatement()
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
		p.Errors.AddSyntaxError("Sau 'cho' phải là một tên định danh", p.curToken)
	}
	stmt.Ident = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	hasAssign := false

	// Nếu có lệnh gán
	if p.peekTokenIs(token.Assign) {
		hasAssign = true
		p.advanceToken()
		p.advanceToken()

		stmt.Value = p.parseExpression(LOWEST)
	}

	// Nếu có mệnh đề 'thuộc'
	if p.peekTokenIs(token.Belong) {
		p.advanceToken()
		p.advanceToken()

		stmt.SetType = p.parseExpression(LOWEST)

	} else if !hasAssign {
		p.Errors.AddSyntaxError("Khai báo biến phải có giá trị khởi tạo hoặc miền xác định", p.curToken)
	}

	return stmt
}

func (p *Parser) parseImplyStatement() *ast.ImplyStatement {
	stmt := &ast.ImplyStatement{Token: p.curToken}
	p.advanceToken()

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{LeftBrace: p.curToken}
	block.Statements = []ast.Statement{}

	if !p.expectPeek(token.LBrace) {
		return nil
	}

	for !p.curTokenIs(token.RBrace) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.advanceToken()
	}

	block.RightBrace = p.curToken
	if !p.expectCur(token.RBrace) {
		return nil
	}

	return block
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{Token: p.curToken}
	stmt.Ident = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	p.advanceToken()
	p.advanceToken()

	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}
