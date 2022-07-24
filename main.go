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
			errors := errorhandler.NewTokenErrorList(input, filepath)

			l := lexer.New(input, errors)
			p := parser.New(l, errors)
			ev := evaluator.New(errors)

			program := p.ParseProgram()

			value := ev.Eval(program)

			if errors.NotEmpty() {
				fmt.Print(errors)

			} else {
				fmt.Println(value.Display())
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

		errors := errorhandler.NewTokenErrorList(line, "")

		l := lexer.New(line, errors)
		p := parser.New(l, errors)
		ev := evaluator.New(errors)

		program := p.ParseProgram()

		value := ev.Eval(program)

		if errors.NotEmpty() {
			fmt.Print(errors)

		} else {
			fmt.Println(value.Display())
		}
	}
}
