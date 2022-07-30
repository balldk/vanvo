package repl

import (
	"bytes"
	"fmt"
	"vila/pkg/errorhandler"
	"vila/pkg/evaluator"
	"vila/pkg/lexer"
	"vila/pkg/object"
	"vila/pkg/parser"

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

		errors := errorhandler.NewErrorList(line, "")

		l := lexer.New(line, errors)
		p := parser.New(l, errors)
		ev := evaluator.New(env, errors)

		program := p.ParseProgram()

		value := ev.Eval(program)

		if errors.NotEmpty() {
			fmt.Print(errors)

		} else {
			fmt.Println(value.Display())
		}
	}
}
