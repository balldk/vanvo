package repl

import (
	"bytes"
	"fmt"
	"strings"
	"vila/pkg/evaluator"
	"vila/pkg/object"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

const PROMPT = ">> "

func welcomeBoard() {
	color.Blue("Chào mừng đến với Vila 0.1.0")
}

func Start() {
	var prompt bytes.Buffer
	color.New(color.FgGreen).Fprint(&prompt, PROMPT)

	rl, err := readline.New(prompt.String())
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	welcomeBoard()

	blockInput := ""
	env := object.NewEnvironment()
	for {
		line, err := rl.Readline()
		line = strings.Trim(line, " ")

		if err != nil {
			fmt.Println("Bái bai :(")
			break
		}

		input := blockInput + line
		lastWord := input[len(input)-1]

		if line == "" {
			blockInput = ""
			rl.SetPrompt(prompt.String())
		}

		if lastWord == ':' || lastWord == '(' {
			blockInput = input + "\n"
			rl.SetPrompt(".. ")
		}

		if blockInput == "" {
			value, errors := evaluator.EvalFromInput(input, "", env)

			if errors.NotEmpty() {
				fmt.Print(errors)

			} else if value != evaluator.NO_PRINT {
				fmt.Println(value.Display())
			}

		} else {
			blockInput = input + "\n"
		}
	}
}
