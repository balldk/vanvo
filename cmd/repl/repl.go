package repl

import (
	"bytes"
	"fmt"
	"vila/pkg/evaluator"
	"vila/pkg/object"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

const PROMPT = ">> "

func Start() {
	var buf bytes.Buffer
	color.New(color.FgGreen).Fprint(&buf, PROMPT)

	rl, err := readline.New(buf.String())
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	env := object.NewEnvironment()
	for {
		line, err := rl.Readline()

		if err != nil {
			fmt.Println("Bái bai :(")
			break
		}

		value, errors := evaluator.EvalFromInput(line, "", env)

		if errors.NotEmpty() {
			fmt.Print(errors)

		} else if value != evaluator.NO_PRINT {
			fmt.Println(value.Display())
		}
	}
}
