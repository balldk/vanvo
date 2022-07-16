package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"vila/errorhandler"
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

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
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

		l := lexer.New(line, lexerErr)
		p := parser.New(l, parserErr)

		res := p.ParseProgram()

		if lexerErr.Length() > 0 {
			fmt.Print(lexerErr)

		} else if parserErr.Length() > 0 {
			fmt.Print(parserErr)
		} else {
			fmt.Print(res)
		}
	}
}
