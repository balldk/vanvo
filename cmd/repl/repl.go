package repl

import (
	"bytes"
	"fmt"
	"strings"
	"vanvo/pkg/evaluator"
	"vanvo/pkg/object"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

const PROMPT = ">> "

func welcomeBoard() {
	color.Blue("Chào mừng đến với VanVo 0.1.0")
	color.Blue(`        _           ?  `)
	color.Blue(`   ┬  ┬┌─┐┌┐┌  ┬  ┬┌─┌'`)
	color.Blue(`   └┐┌┘├─┤│││  └┐┌┘│ │ `)
	color.Blue(`    └┘ ┴ ┴┘└┘   └┘ └─┘ `)
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
		spaces := strings.Repeat(" ", 4)
		line = strings.ReplaceAll(line, "\t", spaces)

		if err != nil {
			fmt.Println("Bái bai :(")
			break
		}

		input := blockInput + line
		lastWord := byte(0)
		if len(line) > 0 {
			lastWord = input[len(input)-1]
		}

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
