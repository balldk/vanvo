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
	p.registerPrefix(token.Ident, p.parseIdentifier)
	p.registerPrefix(token.Int, p.parseInt)
	p.registerPrefix(token.Real, p.parseReal)
	p.registerPrefix(token.True, p.parseBoolean)
	p.registerPrefix(token.False, p.parseBoolean)
	p.registerPrefix(token.Bang, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)
	p.registerPrefix(token.Plus, p.parsePrefixExpression)
	p.registerPrefix(token.LParen, p.parseGroupedExpression)
	p.registerPrefix(token.LBracket, p.parseInterval)
	p.registerPrefix(token.If, p.parseIfExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Minus, p.parseInfixExpression)
	p.registerInfix(token.Slash, p.parseInfixExpression)
	p.registerInfix(token.Asterisk, p.parseInfixExpression)
	p.registerInfix(token.Dot, p.parseInfixExpression)
	p.registerInfix(token.Hat, p.parseInfixExpression)
	p.registerInfix(token.Equal, p.parseInfixExpression)
	p.registerInfix(token.NotEqual, p.parseInfixExpression)
	p.registerInfix(token.Less, p.parseInfixExpression)
	p.registerInfix(token.Greater, p.parseInfixExpression)
	p.registerInfix(token.LessEqual, p.parseInfixExpression)
	p.registerInfix(token.GreaterEqual, p.parseInfixExpression)
	p.registerInfix(token.LParen, p.parseCallExpression)

	p.advanceToken()
	p.advanceToken()

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

func (p *Parser) advanceToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.AdvanceToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.advanceToken()
		if !p.curIsStatementSeperator() {
			p.invalidSyntax()
		}
		p.advanceToken()
	}

	return program
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curIsStatementSeperator() bool {
	return p.curTokenIs(token.Semicolon) || p.curTokenIs(token.Endline) || p.curTokenIs(token.EOF)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.advanceToken()
		return true
	} else {
		p.advanceToken()
		p.expectError(t)
		return false
	}
}

func (p *Parser) expectCur(t token.TokenType) bool {
	if p.curTokenIs(t) {
		p.advanceToken()
		return true
	} else {
		p.expectError(t)
		p.advanceToken()
		return false
	}
}

func (p *Parser) syntaxError(message string) {
	p.Errors.AddParserError(message, p.curToken)
}

func (p *Parser) invalidSyntax() {
	p.Errors.AddParserError("Cú pháp không hợp lệ", p.curToken)
}

func (p *Parser) expectError(tokType token.TokenType) {
	var msg string

	if p.curIsStatementSeperator() {
		msg = fmt.Sprintf("Thiếu `%s`", string(tokType))
	} else {
		msg = fmt.Sprintf("Cần `%s` thay vì `%s`", string(tokType), string(p.curToken.Literal))
	}
	p.syntaxError(msg)
}
