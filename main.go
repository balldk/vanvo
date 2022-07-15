package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"vila/lexer"
	"vila/parser"
	"vila/token"
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
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("Can't read file:", os.Args[1])
		} else {
			l := lexer.New(string(file))

			if l.Errors.Length() > 0 {
				fmt.Println(l.Errors)
				return
			}

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Println(tok)
			}
		}
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)
		if ok := scanner.Scan(); !ok {
			return
		}

		line := scanner.Text()
		// line := "cho abc = 5"
		l := lexer.New(line)
		p := parser.New(l)

		if l.Errors.Length() > 0 {
			fmt.Println("Lexer error:")
			fmt.Println(l.Errors)
			break
		}
		if p.Errors.Length() > 0 {
			fmt.Println("Parser error:")
			fmt.Println(p.Errors)
			break
		}
		fmt.Println(p.ParseProgram())
		// for _, each := range p.ParseProgram() {

		// }
		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	if l.Errors.Length() > 0 {
		// 		fmt.Println(l.Errors)
		// 		break
		// 	}
		// 	fmt.Println(tok)
		// }
	}
}
