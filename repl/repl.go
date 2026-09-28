package repl

import (
	"bufio"
	"fmt"
	"io"

	"bihari/evaluator"
	"bihari/lexer"
	"bihari/object"
	"bihari/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := evaluator.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Allow the user to exit
		if line == "exit" || line == "quit" {
			fmt.Fprintln(out, "Chala bhai mila tani fir!")
			return
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() == object.ERROR {
				fmt.Fprintln(out, evaluated.Inspect())
			} else {
				io.WriteString(out, evaluated.Inspect()+"\n")
			}
		}
	}
}
