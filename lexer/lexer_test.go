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
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.L_PAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.R_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.L_PAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.R_PAREN, ")"},
		{token.SEMICOLON, ""},
		{token.EOF, ""},
	}

	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y
	};

	let result = add(five, ten);
	`
	lexer := NewLexer(input)

	for i, ts := range tokenSet {
		token := lexer.GetToken()

		if token.Type != ts.expectedType {
			t.Errorf("index: %d, expected: %v, got: %v", i, ts.expectedType, token.Type)
		}

		if token.Literal != ts.expectedLiteral {
			t.Errorf("index: %d, expected: %v, got: %v", i, ts.expectedLiteral, token.Literal)
		}
		lexer.ReadChar()
	}
}
