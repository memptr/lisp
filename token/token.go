package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// TODO: enrich
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INT = "INT"

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	SYMBOL = "SYMBOL"

	LPAREN = "("
	RPAREN = ")"
)
