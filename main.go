package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"vila/errorhandler"
	"vila/evaluator"
	"vila/lexer"
	"vila/parser"
	"vila/token"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

const PROMPT = ">> "

func errRecover() {
	if r := recover(); r != nil {
		fmt.Print("")
	}
}

func main() {
	defer errRecover()

	if len(os.Args) > 1 {
		filepath := os.Args[1]
		file, err := ioutil.ReadFile(filepath)
		input := string(file)

		if err != nil {
			fmt.Println("Can't read file:", filepath)
		} else {
			lexerErr := errorhandler.NewErrorList(input, filepath)
			l := lexer.New(input, lexerErr)

			for tok := l.AdvanceToken(); tok.Type != token.EOF; tok = l.AdvanceToken() {
				fmt.Println(tok)
			}
			if lexerErr.Length() > 0 {
				fmt.Print(lexerErr)
				return
			}
		}
		return
	}

	var buf bytes.Buffer
	color.New(color.FgGreen).Fprint(&buf, PROMPT)

	rl, err := readline.New(buf.String())
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()

		if err != nil {
			fmt.Println("Bái bai :(")
			break
		}

		lexerErr := errorhandler.NewErrorList(line, "")
		parserErr := errorhandler.NewErrorList(line, "")
		evaluatorErr := errorhandler.NewErrorList(line, "")

		l := lexer.New(line, lexerErr)
		p := parser.New(l, parserErr)
		ev := evaluator.New(evaluatorErr)

		program := p.ParseProgram()
		value := ev.Eval(program)

		if lexerErr.NotEmpty() {
			fmt.Print(lexerErr)

		} else if parserErr.NotEmpty() {
			fmt.Print(parserErr)

		} else if evaluatorErr.NotEmpty() {
			fmt.Print(evaluatorErr)

		} else {
			fmt.Println(value.Display())
		}
	}
}
