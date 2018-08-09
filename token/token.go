package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers.
	IDENT = "IDENT"

	// Literals.
	INT = "INT"

	// Operators.
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"

	// Groups.
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords.
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
