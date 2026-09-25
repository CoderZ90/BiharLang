package lexer

import (
	"bihari/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `rakha naam = "Babu"
rakha umar = 20
agar umar > 18 {
	bol "Bada ba"
}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.RAKHA, "rakha"},
		{token.IDENT, "naam"},
		{token.ASSIGN, "="},
		{token.STRING, "Babu"},
		{token.RAKHA, "rakha"},
		{token.IDENT, "umar"},
		{token.ASSIGN, "="},
		{token.INT, "20"},
		{token.AGAR, "agar"},
		{token.IDENT, "umar"},
		{token.GT, ">"},
		{token.INT, "18"},
		{token.LBRACE, "{"},
		{token.BOL, "bol"},
		{token.STRING, "Bada ba"},
		{token.RBRACE, "}"},
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
