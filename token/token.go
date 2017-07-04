package token

type TokenType string

type token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL signifies a token/character we don't know
	ILLEGAL = "ILLEGAL"
	/*  EOF stands for gg"end of file", which tells our parser
	    later on that it can stop. */
	EOF = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
