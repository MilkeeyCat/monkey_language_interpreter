package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Milk-Interpreters/milk_interpreter_go/lexer"
	"github.com/Milk-Interpreters/milk_interpreter_go/token"
)

var PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
