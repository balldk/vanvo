package error

import (
	"github.com/fatih/color"
)

type ErrorType string

const (
	SYNTAX_ERROR = "Lỗi cú pháp"
)

func ThrowError(errType ErrorType, message string) {
	red := color.New(color.FgHiRed)
	white := color.New(color.FgWhite)

	red.Printf("\n%v: ", errType)
	white.Println(message)

	panic("")
}
