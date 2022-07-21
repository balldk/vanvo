package errorhandler

import (
	"bytes"
	"strings"
	"vila/token"

	"github.com/fatih/color"
)

type ErrorType string

const (
	SYNTAX_ERROR  = "Lỗi cú pháp"
	RUNTIME_ERROR = "Lỗi"
)

type Error struct {
	Type    ErrorType
	Message string
	Token   token.Token
}

type ErrorList struct {
	filepath string
	lines    []string
	Errors   []Error
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func NewError(errType ErrorType, message string, tok token.Token) Error {
	return Error{Type: errType, Message: message, Token: tok}
}

func NewErrorList(input string, filepath string) *ErrorList {
	return &ErrorList{
		Errors:   []Error{},
		lines:    strings.Split(input, "\n"),
		filepath: filepath,
	}
}

func (eh *ErrorList) AddError(err Error) {
	eh.Errors = append(eh.Errors, err)
}

func (eh *ErrorList) AddSyntaxError(message string, tok token.Token) {
	err := NewError(SYNTAX_ERROR, message, tok)
	eh.AddError(err)
}

func (eh *ErrorList) AddRuntimeError(message string, tok token.Token) {
	err := NewError(RUNTIME_ERROR, message, tok)
	eh.AddError(err)
}

func (eh *ErrorList) Length() int {
	return len(eh.Errors)
}

func (eh *ErrorList) NotEmpty() bool {
	return len(eh.Errors) > 0
}

func (el *ErrorList) String() string {
	var buf bytes.Buffer

	red := color.New(color.FgHiRed)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgHiGreen)
	white := color.New(color.FgWhite)

	if el.filepath != "" {
		green.Fprintln(&buf, "-->", el.filepath)
	}

	for index, err := range el.Errors {

		fromLine := max(1, err.Token.Line-1)
		toLine := min(err.Token.Line, len(el.lines))

		red.Fprint(&buf, err.Type+": ")
		white.Fprintln(&buf, err.Message)

		for i := fromLine - 1; i < toLine; i++ {
			line := el.lines[i]
			if i+1 == err.Token.Line {
				blue.Fprint(&buf, i+1, " | ")
				white.Fprintln(&buf, line)
				blue.Fprint(&buf, " ", " | ")
				white.Fprint(&buf, strings.Repeat(" ", err.Token.Column-1))

				for j := 0; j < max(1, len(err.Token.Literal)); j++ {
					red.Fprint(&buf, "^")
				}
			} else {
				blue.Fprint(&buf, " ", " | ")
				white.Fprintln(&buf, line)
			}
		}
		buf.WriteString("\n")
		if index != len(el.Errors)-1 {
			buf.WriteString("\n")
		}
	}

	return buf.String()
}
