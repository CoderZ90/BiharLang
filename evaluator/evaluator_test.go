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

func TestFunctionObject(t *testing.T) {
	input := "kaam jor(x, y) { lauta x + y }"

	evaluated := testEval(input)
	function, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("Expected *object.Function, got %T (%v)", evaluated, evaluated)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("Expected 2 parameters, got %d", len(function.Parameters))
	}

	expectedParams := []string{"x", "y"}
	for i, param := range function.Parameters {
		if param.Value != expectedParams[i] {
			t.Errorf("Expected parameter %d to be '%s', got '%s'", i, expectedParams[i], param.Value)
		}
	}

	if function.Body == nil {
		t.Fatalf("Function body is nil")
	}

	if len(function.Body.Statements) != 1 {
		t.Fatalf("Expected 1 statement in body, got %d", len(function.Body.Statements))
	}
}

func TestFunctionCalls(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// Identity function
		{"kaam identity(x) { lauta x } identity(5)", 5},

		// Double function
		{"kaam double(x) { lauta x * 2 } double(5)", 10},

		// Add function
		{"kaam add(x, y) { lauta x + y } add(5, 5)", 10},

		// Nested function calls
		{"kaam add(x, y) { lauta x + y } add(5 + 5, add(5, 5))", 20},

		// Complex nested
		{"kaam add(x, y) { lauta x + y } kaam mul(x, y) { lauta x * y } add(mul(2, 3), 4)", 10},

		// Function returning function result
		{"kaam jor(x, y) { lauta x + y } jor(1, 2)", 3},

		// Implicit return (last expression)
		{"kaam test() { 10 } test()", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
	rakha x = 10
	kaam getX() { lauta x }
	getX()
	`

	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 10)
}

func TestFunctionErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		// Calling non-function
		{"rakha x = 10 x()", "arre bhai! 'INTEGER' kaam nahi ba! 💀"},

		// Wrong number of arguments
		{"kaam add(x, y) { lauta x + y } add(1)", "arre bhai! 2 params chahiye, 1 milal 💀"},
		{"kaam add(x, y) { lauta x + y } add(1, 2, 3)", "arre bhai! 2 params chahiye, 3 milal 💀"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("Expected *object.Error, got %T (%v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("Expected error message '%s', got '%s'", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestStringConcatenationInFunctions(t *testing.T) {
	input := `
	kaam greeting(naam) { lauta "Ka ho " + naam + "!" }
	greeting("Babu")
	`

	evaluated := testEval(input)
	strObj, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("Expected *object.String, got %T (%v)", evaluated, evaluated)
	}

	if strObj.Value != "Ka ho Babu!" {
		t.Errorf("Expected 'Ka ho Babu!', got '%s'", strObj.Value)
	}
}

func TestModuloExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"10 % 3", 1},
		{"10 % 2", 0},
		{"7 % 7", 0},
		{"20 % 6", 2},

		// Precedence: % binds tighter than +
		{"10 + 6 % 4", 12},
		{"6 % 4 + 10", 12},

		// Even/odd pattern you'll use constantly
		{"10 % 2 == 0", 1},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		// Handle boolean expectations
		if tt.expected == 1 && tt.input == "10 % 2 == 0" {
			testBooleanObject(t, evaluated, true)
			continue
		}

		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestDivisionByZero(t *testing.T) {
	tests := []string{
		"10 / 0",
		"10 % 0",
	}

	for _, input := range tests {
		evaluated := testEval(input)
		_, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("Expected *object.Error for %s, got %T", input, evaluated)
		}
	}
}
