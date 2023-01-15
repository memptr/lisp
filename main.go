package main

import (
	"lisp/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
