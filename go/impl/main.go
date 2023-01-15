package main

import (
	"impl/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
