package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Identifiers and literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LBRACE   = "{"
	RBRACE   = "}"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	MACRO    = "MACRO"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Literal string
	Type    TokenType
}

type TokenType string

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}
