package main

import (
	"fmt"
	"os"

	"bihari/repl"
)

func main() {
	fmt.Println(`
	█████▄ ▄▄ ▄▄ ▄▄  ▄▄▄  ▄▄▄▄  ██      ▄▄▄  ▄▄  ▄▄  ▄▄▄▄  
	██▄▄██ ██ ██▄██ ██▀██ ██▄█▄ ██     ██▀██ ███▄██ ██ ▄▄  
	██▄▄█▀ ██ ██ ██ ██▀██ ██ ██ ██████ ██▀██ ██ ▀██ ▀███▀  

	A Toy Programming Language, In Bhojpuri :).              
	v0.0.0                           
	`)

	fmt.Println()
	fmt.Println("        Kaa re bhai, tayyar baa? Chala code kari :)")
	fmt.Println("        Type 'exit' to quit.")
	fmt.Println()
	fmt.Println("──────────────────────────────────────────────────────────────")
	fmt.Println()

	repl.Start(os.Stdin, os.Stdout)
}
