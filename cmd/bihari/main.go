package main

import (
	"fmt"
	"os"

	"bihari/repl"
)

func main() {
	fmt.Println("Bihari Programming Language")
	fmt.Println("Kaa reee ka haal baa bhai.")
	fmt.Println("Type Bihari code below. Type 'exit' to quit.")
	fmt.Println("----------------------------------------")
	repl.Start(os.Stdin, os.Stdout)
}
