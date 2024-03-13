package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MilkeeyCat/monkey_language_interpreter/evaluator"
	"github.com/MilkeeyCat/monkey_language_interpreter/lexer"
	"github.com/MilkeeyCat/monkey_language_interpreter/object"
	"github.com/MilkeeyCat/monkey_language_interpreter/parser"
)

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

var PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
