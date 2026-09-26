package parser

import (
	"bihari/ast"
	"bihari/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `rakha x = 10`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	stmt := program.Statements[0]

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected *ast.LetStatement, got %T", stmt)
		return
	}

	if letStmt.Name.Value != "x" {
		t.Errorf("Expected variable name 'x', got '%s'", letStmt.Name.Value)
	}
}
