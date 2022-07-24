package errorhandler

import (
	"bytes"
	"strings"
	"vila/ast"
	"vila/token"

	"github.com/fatih/color"
)

type ErrorType string

const (
	SYNTAX_ERROR  = "Lỗi cú pháp"
	RUNTIME_ERROR = "Lỗi"
)

var (
	red   = color.New(color.FgHiRed)
	blue  = color.New(color.FgBlue)
	green = color.New(color.FgHiGreen)
	white = color.New(color.FgWhite)
)

type TokenError struct {
	Type    ErrorType
	Message string
	Token   token.Token
}

func NewTokenError(errType ErrorType, message string, tok token.Token) TokenError {
	return TokenError{Type: errType, Message: message, Token: tok}
}

type NodeError struct {
	Type    ErrorType
	Message string
	Node    ast.Node
}

func NewNodeError(errType ErrorType, message string, node ast.Node) NodeError {
	return NodeError{Type: errType, Message: message, Node: node}
}

type ErrorList struct {
	filepath     string
	lines        []string
	LexerErrors  []TokenError
	ParserErrors []TokenError
	EvalErrors   []NodeError
}

func NewTokenErrorList(input string, filepath string) *ErrorList {
	return &ErrorList{
		LexerErrors:  []TokenError{},
		ParserErrors: []TokenError{},
		EvalErrors:   []NodeError{},
		lines:        strings.Split(input, "\n"),
		filepath:     filepath,
	}
}

func (eh *ErrorList) AddLexerError(message string, tok token.Token) {
	err := NewTokenError(SYNTAX_ERROR, message, tok)
	eh.LexerErrors = append(eh.LexerErrors, err)
}

func (eh *ErrorList) AddParserError(message string, tok token.Token) {
	err := NewTokenError(SYNTAX_ERROR, message, tok)
	eh.ParserErrors = append(eh.ParserErrors, err)
}

func (eh *ErrorList) AddRuntimeError(message string, node ast.Node) {
	err := NewNodeError(RUNTIME_ERROR, message, node)
	eh.EvalErrors = append(eh.EvalErrors, err)
}

func (eh *ErrorList) NotEmpty() bool {
	return len(eh.LexerErrors) > 0 || len(eh.ParserErrors) > 0 || len(eh.EvalErrors) > 0
}

func (el *ErrorList) printLine(buf *bytes.Buffer, lineNum int, showLine bool) {
	line := el.lines[lineNum-1]
	if showLine {
		blue.Fprint(buf, lineNum, " | ")
	} else {
		blue.Fprint(buf, " ", " | ")
	}
	white.Fprintln(buf, line)
}

func (el *ErrorList) underline(buf *bytes.Buffer, from, length int) {
	blue.Fprint(buf, " ", " | ")
	white.Fprint(buf, strings.Repeat(" ", from))

	for i := 0; i < length; i++ {
		red.Fprint(buf, "^")
	}
}

func (el *ErrorList) printErrorMessage(buf *bytes.Buffer, t ErrorType, message string) {
	red.Fprint(buf, t+": ")
	white.Fprintln(buf, message)
}

func (el *ErrorList) printTokenErrors(buf *bytes.Buffer, errors []TokenError) {
	for index, err := range errors {
		fromLine := max(1, err.Token.Line-1)
		toLine := min(err.Token.Line, len(el.lines))

		el.printErrorMessage(buf, err.Type, err.Message)

		for i := fromLine - 1; i < toLine; i++ {
			if i+1 == err.Token.Line {
				el.printLine(buf, i+1, true)
				el.underline(buf, err.Token.Column-1, max(1, len(err.Token.Literal)))
			} else {
				el.printLine(buf, i+1, false)
			}
		}
		buf.WriteString("\n")
		if index != len(errors)-1 {
			buf.WriteString("\n")
		}
	}
}

func (el *ErrorList) printEvalErrors(buf *bytes.Buffer) {
	for index, err := range el.EvalErrors {
		node := el.EvalErrors[index].Node
		fromTok := node.FromToken()
		toTok := node.ToToken()

		fromLine := max(1, fromTok.Line-1)
		toLine := min(toTok.Line, len(el.lines))

		el.printErrorMessage(buf, err.Type, err.Message)

		for i := fromLine - 1; i < toLine; i++ {
			line := el.lines[i]

			if i+1 == fromTok.Line {
				el.printLine(buf, i+1, true)
				if fromTok.Line != toTok.Line {
					el.underline(buf, fromTok.Column-1, len(line)-i+1)
				} else {
					el.underline(buf, fromTok.Column-1, toTok.Column-fromTok.Column+len(toTok.Literal))
				}

			} else if i+1 == toTok.Line {
				el.printLine(buf, i+1, true)
				el.underline(buf, 0, toTok.Column+len(toTok.Literal))

			} else if fromTok.Line < i+1 && i+1 < toTok.Line {
				el.printLine(buf, i+1, true)
				el.underline(buf, 0, len(line))
			} else {
				el.printLine(buf, i+1, false)
			}
		}
		buf.WriteString("\n")
		if index != len(el.EvalErrors)-1 {
			buf.WriteString("\n")
		}
	}
}

func (el *ErrorList) String() string {
	var buf bytes.Buffer
	if !el.NotEmpty() {
		return ""
	}

	if el.filepath != "" {
		green.Fprintln(&buf, "-->", el.filepath, "\n")
	}

	if len(el.LexerErrors) > 0 {
		el.printTokenErrors(&buf, el.LexerErrors)

	} else if len(el.ParserErrors) > 0 {
		el.printTokenErrors(&buf, el.ParserErrors)

	} else if len(el.EvalErrors) > 0 {
		el.printEvalErrors(&buf)
	}

	return buf.String()
}
