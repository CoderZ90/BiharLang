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
	return Eval(program)
}

func TestEvalIntegerLiteral(t *testing.T) {
	// This test will FAIL for now because Eval doesn't handle Program yet.
	input := `10`
	evaluated := testEval(input)

	intObj, ok := evaluated.(*object.Integer)
	if !ok {
		t.Fatalf("Expected *object.Integer, got %T", evaluated)
	}

	if intObj.Value != 10 {
		t.Fatalf("Expected 10, got %d", intObj.Value)
	}
}
