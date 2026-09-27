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

	valStmt, ok := letStmt.Value.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("Expected *ast.IntegerLiteral, got %T", letStmt.Value)
		return
	}

	if valStmt.Value != 10 {
		t.Errorf("Expected value 10, got %d", valStmt.Value)
	}
}

func TestStringStatements(t *testing.T) {
	input := `rakha naam = "Babu"`

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
		t.Fatalf("Expected *ast.LetStatement, got %T", stmt)
	}

	if letStmt.Name.Value != "naam" {
		t.Fatalf("Expected variable name 'naam', got '%s'", letStmt.Name.Value)
	}

	strLit, ok := letStmt.Value.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("Expected *ast.StringLiteral, got %T", letStmt.Value)
	}

	if strLit.Value != "Babu" {
		t.Fatalf("Expected value 'Babu', got '%s'", strLit.Value)
	}
}

func TestPrintStatements(t *testing.T) {
	input := `bol "Ka ho duniya!"`

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
	printStmt, ok := stmt.(*ast.PrintStatement)
	if !ok {
		t.Fatalf("Expected *ast.PrintStatement, got %T", stmt)
	}

	strLit, ok := printStmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("Expected *ast.StringLiteral, got %T", printStmt.Expression)
	}

	if strLit.Value != "Ka ho duniya!" {
		t.Fatalf("Expected 'Ka ho duniya!', got '%s'", strLit.Value)
	}
}
