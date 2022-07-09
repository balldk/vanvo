package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"vila/lexer"
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
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Println(tok)
		}
	}
}
