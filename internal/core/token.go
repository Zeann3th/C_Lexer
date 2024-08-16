package core

type Token int

const (
	_ = iota
	TOKEN_HASH
	TOKEN_COMMENT
	TOKEN_SYMBOL
	TOKEN_OPAREN
	TOKEN_CPAREN
	TOKEN_OCURLY
	TOKEN_CCURLY
	TOKEN_SEMICOLON
	TOKEN_NUMBER
	TOKEN_STRING
	TOKEN_CHAR
	TOKEN_RETURN
	TOKEN_ERROR
	TOKEN_EOF
)
