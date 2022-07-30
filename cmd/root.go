package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"vila/cmd/repl"
	"vila/pkg/evaluator"
	"vila/pkg/object"
)

func errRecover() {
	if r := recover(); r != nil {
		fmt.Print("Lỗi hệ thống")
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
		env := object.NewEnvironment()

		value, errors := evaluator.EvalFromInput(input, path, env)

		if errors.NotEmpty() {
			fmt.Print(errors)

		} else if value == evaluator.NO_PRINT {
			fmt.Println(value.Display())
		}
	}
}

func Execute() {
	defer errRecover()
	initConfig()

	if len(os.Args) > 1 {
		runFromFile()

	} else {
		repl.Start()
	}

}
