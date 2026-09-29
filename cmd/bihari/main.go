package main

import (
	"fmt"
	"os"

	"bihari/evaluator"
	"bihari/lexer"
	"bihari/object"
	"bihari/parser"
	"bihari/repl"
)

const (
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Reset  = "\033[0m"
)

func main() {
	printBanner()

	if len(os.Args) > 1 {
		runFile(os.Args[1])
	} else {
		runREPL()
	}
}

func printBanner() {
	fmt.Println(`
	█████▄ ▄▄ ▄▄ ▄▄  ▄▄▄  ▄▄▄▄  ██      ▄▄▄  ▄▄  ▄▄  ▄▄▄▄  
	██▄▄██ ██ ██▄██ ██▀██ ██▄█▄ ██     ██▀██ ███▄██ ██ ▄▄  
	██▄▄█▀ ██ ██ ██ ██▀██ ██ ██ ██████ ██▀██ ██ ▀██ ▀███▀  

	A Toy Programming Language, In Bhojpuri :).              
	v0.0.0
	`)

	fmt.Println()
}

func runREPL() {
	fmt.Println()
	fmt.Println("        Kaa re bhai, tayyar baa? Chala code kari :)")
	fmt.Println("        Type 'exit' to quit.")
	fmt.Println()
	fmt.Println("───────────────────────────────────────────────────────────────")
	fmt.Println()

	repl.Start(os.Stdin, os.Stdout)
}

func runFile(filename string) {
	fmt.Printf(Yellow+"Compiler Running: %s\n"+Reset, filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Arre bhaiya! '%s' file kidhar baa! Error: %s\n", filename, err)
		os.Exit(1)
	}

	input := string(data)

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		for _, msg := range p.Errors() {
			fmt.Println("❌ Syntax Error:", msg)
		}
		os.Exit(1)
	}

	env := evaluator.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	if evaluated != nil && evaluated.Type() == object.ERROR {
		fmt.Println("❌ Runtime Error:", evaluated.Inspect())
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(Green + "Bahut Badhiya bhai 🎉" + Reset)
}
