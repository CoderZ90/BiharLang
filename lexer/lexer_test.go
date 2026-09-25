package lexer

import (
	"bihari/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `rakha x = 10`

	// The tokens we expect the lexer to produce, in exact order
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.RAKHA, "rakha"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
