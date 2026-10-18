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

	if len(os.Args) < 2 {
		runREPL()
		return
	}

	switch os.Args[1] {
	case "compile", "run", "c", "r":
		if len(os.Args) < 3 {
			fmt.Println(Red + "💀 Arre bhai! File naam ta batao! Example: bh compile main.bh" + Reset)
			os.Exit(1)
		}
		runFile(os.Args[2])

	case "repl":
		runREPL()

	case "help", "--help", "-h":
		printHelp()

	default:
		runFile(os.Args[1])
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

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  bh <file.bh>          Run a Bihari file")
	fmt.Println("  bh compile <file.bh>  Compile and run a Bihari file")
	fmt.Println("  bh run <file.bh>      Same as compile")
	fmt.Println("  bh repl               Start interactive REPL")
	fmt.Println("  bh help               Show this help message")
	fmt.Println()
	fmt.Println("Kaa re bhai, samajh aail?!")
}

func runREPL() {
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
		fmt.Printf(Red+"Arre bhaiya! '%s' file kidhar baa! Error: %s\n"+Reset, filename, err)
		os.Exit(1)
	}

	input := string(data)

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		for _, msg := range p.Errors() {
			fmt.Println(Red + "❌ Syntax Error: " + msg + Reset)
		}
		os.Exit(1)
	}

	env := evaluator.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	if evaluated != nil && evaluated.Type() == object.ERROR {
		fmt.Println(Red + "❌ Runtime Error: " + evaluated.Inspect() + Reset)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(Green + "Bahut Badhiya bhai 🎉" + Reset)
}
