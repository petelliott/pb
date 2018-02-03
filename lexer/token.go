package lexer

const (
	IDENTIFIER  = iota
	LITERAL     = iota
	STR_LITERAL = iota
	L_PAREN     = iota
	R_PAREN     = iota
	L_BRACKET   = iota
	R_BRACKET   = iota
	L_BRACE     = iota
	R_BRACE     = iota
	SEMICOLON   = iota
	OPERATOR    = iota
	KEYWORD     = iota
	TYPE        = iota
)

type Token struct {
	tok   int
	value string
}