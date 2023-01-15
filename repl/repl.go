package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			fmt.Println("bye")
			break
		}
		l := lexer.NewLexer(line)

		for t := l.GetToken(); t.Type != token.EOF; t = l.GetToken() {
			fmt.Println(t)
		}
	}
}
