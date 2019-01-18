package token

type TokenType string

type Token struct{
	// Need a type to distinguish b/w different types
	Type TokenType
	// Holds the value
	Literal string
}

// Possible tokens

const(
	ILLEGAL = "ILLEGAL" // token we dont know about
	EOF = "EOF" // Tell parser to stop

	// Identifiers and literals
	IDENT = "IDENT" // add, x, y, foobar, ...
	INT = "INT" // 1234567890

	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	// Delimeters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	IF = "IF"
	ELSE = "ELSE"
	TRUE = "TRUE"
	FALSE = "FALSE"
	RETURN = "RETURN"


)

var keywords = map[string]TokenType{
	"fn" :FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}


func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}