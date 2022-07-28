package parser

import (
	"vila/pkg/ast"
	"vila/pkg/token"
)

func (p *Parser) parseStatement() ast.Statement {
	p.updateIdentLevel()

	var stmt ast.Statement

	if p.curTokenIs(token.Ident) && p.peekTokenIs(token.Assign) {
		stmt = p.parseAssignStatement()
		p.checkEndStatement()
		p.updateIdentLevel()

		return stmt
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

func (p *Parser) parseExpressionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{}
	stmt.Expression = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseLetStatement() ast.Statement {
	letToken := p.curToken

	// Identifier
	if !p.expectPeek(token.Ident) {
		p.Errors.AddParserError("Sau `cho` phải là một tên định danh", p.curToken)
	}
	ident := &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}

	// If declare function
	if p.peekTokenIs(token.LParen) {
		p.advanceToken()
		return p.parseFunction(letToken, ident)
	}

	// Else declare variable
	stmt := &ast.VarDeclareStatement{Token: letToken, Ident: ident}

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

func (p *Parser) parseFunction(letToken token.Token, ident *ast.Identifier) *ast.FunctionDeclareStatement {
	fn := &ast.FunctionDeclareStatement{Token: letToken, Ident: ident}

	for p.peekTokenIs(token.Ident) {
		p.advanceToken()
		param := p.parseIdentifier().(*ast.Identifier)
		fn.Params = append(fn.Params, param)

		if p.peekTokenIs(token.RParen) {
			p.advanceToken()
			break
		}
		if !p.expectPeek(token.Comma) {
			return nil
		}
	}
	if !p.expectCur(token.RParen) {
		return nil
	}

	if !p.expectCur(token.Assign) {
		return nil
	}

	fn.Body = p.parseExpression(LOWEST)

	return fn
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

	if p.curTokenIs(token.Else) {
		stmt.Alternative = p.parseBlockStatement()
	}

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

	if p.curTokenIs(token.Endline) && len(p.curToken.Literal)/4 < curLevel {
		p.skipEndline()
		p.invalidIndent()
	}
	if p.curTokenIs(token.EOF) {
		p.syntaxError("Thiếu mệnh đề sau điều kiện")
	}

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