package parser

import (
	"fmt"
	"vila/ast"
	"vila/errorhandler"
	"vila/lexer"
	"vila/token"
)

func New(l *lexer.Lexer, errors *errorhandler.ErrorList) *Parser {
	p := &Parser{
		l:      l,
		Errors: errors,
	}
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseInt)
	p.registerPrefix(token.REAL, p.parseReal)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.HAT, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NEQ, p.parseInfixExpression)
	p.registerInfix(token.LESS, p.parseInfixExpression)
	p.registerInfix(token.GREATER, p.parseInfixExpression)
	p.registerInfix(token.LESS_EQ, p.parseInfixExpression)
	p.registerInfix(token.GREATER_EQ, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

type Parser struct {
	l      *lexer.Lexer
	Errors *errorhandler.ErrorList

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
		if !p.curIsStatementSeperator() {
			p.syntaxError("Cú pháp không hợp lệ")
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekIsStatementSeperator() bool {
	return p.peekTokenIs(token.SEMICOLON) || p.peekTokenIs(token.ENDLINE) || p.peekTokenIs(token.EOF)
}

func (p *Parser) curIsStatementSeperator() bool {
	return p.curTokenIs(token.SEMICOLON) || p.curTokenIs(token.ENDLINE) || p.curTokenIs(token.EOF)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.nextToken()
		p.expectError(t)
		return false
	}
}

func (p *Parser) syntaxError(message string) {
	p.Errors.AddSyntaxError(message, p.curToken)
}

func (p *Parser) expectError(tokType token.TokenType) {
	var msg string

	if p.curIsStatementSeperator() {
		msg = fmt.Sprintf("Thiếu `%s` ở đây", string(tokType))
	} else {
		msg = fmt.Sprintf("Cần `%s` thay vì `%s` ở đây", string(tokType), string(p.curToken.Literal))
	}
	p.syntaxError(msg)
}
