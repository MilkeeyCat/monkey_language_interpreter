package main

import (
	"fmt"
	"os"

	"github.com/Milk-Interpreters/milk_interpreter_go/repl"
)

func main() {
	fmt.Println("Welcome to milk programming lang!")
	repl.Start(os.Stdin, os.Stdout)
}
