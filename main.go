package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"vila/errorhandler"
	"vila/evaluator"
	"vila/lexer"
	"vila/object"
	"vila/parser"
	"vila/repl"
)

func errRecover() {
	if r := recover(); r != nil {
		fmt.Print("")
	}
}

func main() {
	// defer errRecover()

	if len(os.Args) > 1 {
		filepath := os.Args[1]
		file, err := ioutil.ReadFile(filepath)
		input := string(file)

		if err != nil {
			fmt.Println("Can't read file:", filepath)
		} else {
			errors := errorhandler.NewErrorList(input, filepath)
			env := object.NewEnvironment()

			l := lexer.New(input, errors)
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
	} else {
		repl.Start()
	}
}
