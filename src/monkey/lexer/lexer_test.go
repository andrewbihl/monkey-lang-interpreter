package lexer

import (
	"lexer"
	"testing"

	"monkey/token"
)


func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for i, test := range tests {
		token := l.NextToken()

		if token.Type != test.expectedType {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, token.Literal)
		}
	}
}