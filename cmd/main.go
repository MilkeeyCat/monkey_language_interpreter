package main

import (
	"fmt"
	"os"

	"github.com/MilkeeyCat/monkey_language_interpreter/repl"
)

func main() {
	fmt.Println("Welcome to monkey programming lang!")
	repl.Start(os.Stdin, os.Stdout)
}
