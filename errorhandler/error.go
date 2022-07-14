package errorhandler

import (
	"bytes"
	"vila/token"

	"github.com/fatih/color"
)

type ErrorType string

const (
	SYNTAX_ERROR = "Lỗi cú pháp"
)

type Error struct {
	Type    ErrorType
	Message string
	Line    int
	Row     int
}

type ErrorList struct {
	Errors []Error
}

func NewError(errType ErrorType, message string, line int, row int) Error {
	return Error{Type: errType, Message: message, Line: line, Row: row}
}

func (e *Error) String() string {
	var buf bytes.Buffer

	red := color.New(color.FgHiRed).Add(color.Bold)
	white := color.New(color.FgWhite)

	red.Fprint(&buf, e.Type+": ")
	white.Fprintln(&buf, e.Message)

	return buf.String()
}

func NewErrorList() *ErrorList {
	return &ErrorList{Errors: []Error{}}
}

func (eh *ErrorList) AddError(err Error) {
	eh.Errors = append(eh.Errors, err)
}

func (eh *ErrorList) AddSyntaxError(message string, tok *token.Token) {
	err := NewError(SYNTAX_ERROR, message, tok.Line, tok.Row)
	eh.AddError(err)
}

func (eh *ErrorList) Length() int {
	return len(eh.Errors)
}

func (eh *ErrorList) String() string {
	var buf bytes.Buffer

	for _, err := range eh.Errors {
		buf.WriteString(err.String())
	}

	return buf.String()
}
