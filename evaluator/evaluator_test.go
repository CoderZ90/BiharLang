package evaluator

import (
	"bihari/lexer"
	"bihari/object"
	"bihari/parser"
	"testing"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := NewEnvironment()
	return Eval(program, env)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"rakha x = 10", 10},
		{"rakha x = 50", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		intObj, ok := evaluated.(*object.Integer)
		if !ok {
			t.Fatalf("Expected *object.Integer, got %T (%v)", evaluated, evaluated)
		}
		if intObj.Value != tt.expected {
			t.Fatalf("Expected %d, got %d", tt.expected, intObj.Value)
		}
	}
}

func TestEvalStringExpression(t *testing.T) {
	evaluated := testEval(`rakha naam = "Babu"`)

	strObj, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("Expected *object.String, got %T", evaluated)
	}

	if strObj.Value != "Babu" {
		t.Fatalf("Expected 'Babu', got '%s'", strObj.Value)
	}
}
