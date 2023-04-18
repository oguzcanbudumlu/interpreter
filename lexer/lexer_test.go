package lexer

import (
	"fmt"
	"interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{expectedType: token.EOF, expectedLiteral: ""},
	}

	l := New(`=+(){},;`)

	for i, tt := range tests {
		tok := l.NextNoken()
		fmt.Println(i)
		fmt.Println(tt, tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tt.expectedLiteral)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
