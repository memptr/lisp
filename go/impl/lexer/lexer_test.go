package lexer

import (
	"impl/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	t.Parallel()

	type test struct {
		input    string
		expected []token.Token
	}

	tests := []test{
		{input: "(+ 1 2)", expected: []token.Token{
			{Type: token.LPAREN, Literal: "("},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.INT, Literal: "1"},
			{Type: token.INT, Literal: "2"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.EOF, Literal: ""},
		}},
		{input: `
				(
					(define r 10)
					(define pi 314)
					(* pi (* r r))
				)
			`, expected: []token.Token{
			{Type: token.LPAREN, Literal: "("},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.SYMBOL, Literal: "define"},
			{Type: token.SYMBOL, Literal: "r"},
			{Type: token.INT, Literal: "10"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.SYMBOL, Literal: "define"},
			{Type: token.SYMBOL, Literal: "pi"},
			{Type: token.INT, Literal: "314"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.ASTERISK, Literal: "*"},
			{Type: token.SYMBOL, Literal: "pi"},
			{Type: token.LPAREN, Literal: "("},
			{Type: token.ASTERISK, Literal: "*"},
			{Type: token.SYMBOL, Literal: "r"},
			{Type: token.SYMBOL, Literal: "r"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.RPAREN, Literal: ")"},
			{Type: token.EOF, Literal: ""},
		}},
	}

	for i, tc := range tests {
		l := New(tc.input)

		for j, tt := range tc.expected {
			tok := l.NextToken()

			if tok.Type != tt.Type {
				t.Fatalf("tests[%d][%d] - tokentype wrong. expected=%q, got=%q", i, j, tt.Type, tok.Type)
			}

			if tok.Literal != tt.Literal {
				t.Fatalf("tests[%d][%d] - literal wrong. expected=%q, got=%q", i, j, tt.Literal, tok.Literal)
			}
		}
	}
}
