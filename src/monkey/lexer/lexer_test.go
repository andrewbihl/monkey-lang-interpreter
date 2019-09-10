package lexer

import (
	"fmt"
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {

	type output struct {
		tokenType token.TokenType
		literal   string
	}

	runTest := func(t *testing.T, sourceCode string, expectedOutputs []output) {
		l := New(sourceCode)

		for i, expected := range expectedOutputs {
			fmt.Println("> " + string(expected.tokenType))
			tok := l.NextToken()

			if tok.Type != expected.tokenType {
				t.Fatalf("tests[%d] - wrong tokentype. expected=%q, got=%q",
					i, expected.tokenType, tok.Type)
			}

			if tok.Literal != expected.literal {
				t.Fatalf("tests[%d] - wrong literal. expected=%q, got=%q",
					i, expected.literal, tok.Literal)
			}
		}
	}

	input1 := `let five = 5;
let ten = 10;
   let add = fn(x, y) {
     x + y;
};
   let result = add(five, ten);
   `

	tests1 := []output{
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
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	t.Run("lexer test", func(t *testing.T) {
		runTest(t, input1, tests1)
	})

	input2 := `!-/*5;
   5 < 10 > 5567;
   `
	tests2 := []output{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.STAR, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5567"},
		{token.SEMICOLON, ";"},
	}

	t.Run("lexer test extended", func(t *testing.T) {
		runTest(t, input2, tests2)
	})

	t.Run("lexer test extended", func(t *testing.T) {

	})

}
