package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"vila/cmd/repl"
	"vila/pkg/evaluator"
	"vila/pkg/object"
)

func errRecover() {
	if r := recover(); r != nil {
		fmt.Print("Lỗi trình thông dịch")
	}
}

func runFromFile() {
	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println("đường dẫn không hợp lệ")
	}

	file, err := ioutil.ReadFile(path)
	input := string(file)
	spaces := strings.Repeat(" ", 4)
	input = strings.ReplaceAll(input, "\t", spaces)

	if err != nil {
		fmt.Printf("Không thể mở file: '%s'\n", path)
	} else {
		env := object.NewEnvironment()

		_, errors := evaluator.EvalFromInput(input, path, env)

		if errors.NotEmpty() {
			fmt.Print(errors)

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
