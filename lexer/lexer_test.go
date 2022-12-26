package lexer

import (
	"monkey/token"
	"testing"
)

func TestGetToken(t *testing.T) {
	tokenSet := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.L_PAREN, "("},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.R_BRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	input := "=+(){},;"
	lexer := NewLexer(input)

	for _, ts := range tokenSet {
		token := lexer.GetToken()

		if token.Type != ts.expectedType {
			t.Errorf("expected: %v, got: %v", ts.expectedType, token.Type)
		}

		if token.Literal != ts.expectedLiteral {
			t.Errorf("expected: %v, got: %v", ts.expectedLiteral, token.Literal)
		}
		lexer.ReadChar()
	}
}
