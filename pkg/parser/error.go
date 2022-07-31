package parser

import (
	"fmt"
	"vila/pkg/token"
)

func (p *Parser) syntaxError(message string) {
	p.Errors.AddParserError(message, p.curToken)
}

func (p *Parser) syntaxErrorImportant(message string) {
	p.Errors.AddParserErrorImportant(message, p.curToken)
}

func (p *Parser) invalidSyntax() {
	p.Errors.AddParserError("Cú pháp không hợp lệ", p.curToken)
}

func (p *Parser) invalidIndent() {
	p.syntaxError("Thụt dòng không hợp lệ")
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
