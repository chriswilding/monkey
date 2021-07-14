package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Literal string
	Type    TokenType
}

type TokenType string
