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
		{"10", 10},
		{"50", 50},

		{"10 + 20", 30},
		{"50 - 10", 40},
		{"5 * 5", 25},
		{"10 / 2", 5},

		{"2 + 3 * 4", 14},
		{"2 * 3 + 4", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalStringExpression(t *testing.T) {
	input := `"Babu"`
	evaluated := testEval(input)

	strObj, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("Expected *object.String, got %T", evaluated)
	}
	if strObj.Value != "Babu" {
		t.Fatalf("Expected 'Babu', got '%s'", strObj.Value)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sach", true},
		{"jhoot", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}
	return true
}

func TestComparisonExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"10 > 5", true},
		{"5 > 10", false},
		{"10 < 20", true},
		{"20 < 10", false},
		{"10 == 10", true},
		{"10 != 10", false},
		{"10 != 5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{} // Can be int64 or *object.Null
	}{
		{"agar (sach) { 10 }", int64(10)},
		{"agar (jhoot) { 10 }", NULL},
		{"agar (1 < 2) { 10 }", int64(10)},
		{"agar (1 > 2) { 10 }", NULL},
		{"agar (1 > 2) { 10 } nata { 20 }", int64(20)},
		{"agar (1 < 2) { 10 } nata { 20 }", int64(10)},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		// Check if we expected NULL
		if tt.expected == NULL {
			if evaluated != NULL {
				t.Errorf("Expected NULL, got %T (%v)", evaluated, evaluated)
			}
			continue
		}

		// Otherwise, check for integer
		integer, ok := evaluated.(*object.Integer)
		if !ok {
			t.Errorf("Expected *object.Integer, got %T (%v)", evaluated, evaluated)
			continue
		}

		if integer.Value != tt.expected.(int64) {
			t.Errorf("Expected %d, got %d", tt.expected, integer.Value)
		}
	}
}
