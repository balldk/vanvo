package errorhandler

import (
	"bytes"
	"strings"
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
	filepath string
	input    string
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

func NewError(errType ErrorType, message string, line int, row int) Error {
	return Error{Type: errType, Message: message, Line: line, Row: row}
}

func NewErrorList(input string, filepath string) *ErrorList {
	return &ErrorList{Errors: []Error{}, input: input, filepath: filepath}
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

func (el *ErrorList) String() string {
	var buf bytes.Buffer

	red := color.New(color.FgHiRed)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgHiGreen)
	white := color.New(color.FgWhite)

	lines := strings.Split(el.input, "\n")

	if el.filepath != "" {
		green.Fprintln(&buf, "-->", el.filepath)
	}

	for _, err := range el.Errors {

		fromLine := max(1, err.Line-1)
		toLine := min(err.Line, len(lines))

		red.Fprint(&buf, err.Type+": ")
		white.Fprintln(&buf, err.Message)

		for i := fromLine - 1; i < toLine; i++ {
			line := lines[i]
			if i+1 == err.Line {
				blue.Fprint(&buf, i+1, " | ")
				white.Fprintln(&buf, line)
				blue.Fprint(&buf, " ", " | ")
				white.Fprint(&buf, strings.Repeat(" ", err.Row-1))
				red.Fprintln(&buf, "^")
			} else {
				blue.Fprint(&buf, " ", " | ")
				white.Fprintln(&buf, line)
			}
		}
		buf.WriteString("\n")
	}

	return buf.String()
}
