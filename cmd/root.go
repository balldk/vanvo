package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"vila/cmd/repl"
	"vila/pkg/errorhandler"
	"vila/pkg/evaluator"
	"vila/pkg/lexer"
	"vila/pkg/object"
	"vila/pkg/parser"
)

func errRecover() {
	if r := recover(); r != nil {
		fmt.Print("")
	}
}

func runFromFile() {
	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println("đường dẫn không hợp lệ")
	}

	file, err := ioutil.ReadFile(path)
	input := string(file)

	if err != nil {
		fmt.Printf("Không thể mở file: '%s'\n", path)
	} else {
		errors := errorhandler.NewErrorList(input, path)
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
}

func Execute() {
	// defer errRecover()
	initConfig()

	if len(os.Args) > 1 {
		runFromFile()

	} else {
		repl.Start()
	}

}
