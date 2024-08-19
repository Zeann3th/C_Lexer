package lexer

type TokenKind int

const (
	UNDEFINED TokenKind = iota
	PREPROCESSOR
	COMMENT
	KEYWORD
	TYPE
	SYMBOL
	OPAREN
	CPAREN
	OCURLY
	CCURLY
	NUMBER
	STRING
	CHAR
	RETURN
	EOF
	ERROR
	// Logical operations
	ASSIGN
	NOT
	NOTEQ
	EQ
	LT
	GT
	LTE
	GTE
	ADD
	SUB
	MUL
	DIV
	// Punctuation
	SEMICOLON
	COMMA
	DOT
)

var Codex = map[TokenKind]string{
	PREPROCESSOR: "PREPROCESSOR",
	COMMENT:      "COMMENT",
	KEYWORD:      "KEYWORD",
	TYPE:         "TYPE",
	SYMBOL:       "SYMBOL",
	OPAREN:       "OPAREN",
	CPAREN:       "CPAREN",
	OCURLY:       "OCURLY",
	CCURLY:       "CCURLY",
	SEMICOLON:    "SEMICOLON",
	COMMA:        "COMMA",
	DOT:          "DOT",
	NUMBER:       "NUMBER",
	STRING:       "STRING",
	CHAR:         "CHAR",
	RETURN:       "RETURN",
	ERROR:        "ERROR",
	EOF:          "EOF",
	UNDEFINED:    "UNDEFINED",
	NOT:          "NOT",
	NOTEQ:        "NOTEQUAL",
	EQ:           "EQUAL",
	LT:           "LESS",
	GT:           "GREATER",
	LTE:          "LESS THAN OR EQUAL TO",
	GTE:          "GREATER THAN OR EQUAL TO",
	ADD:          "ADDITION",
	SUB:          "SUBTRACTION",
	DIV:          "DIVISION",
	MUL:          "MULTIPLICATION",
}

type Token struct {
	Name  string
	Kind  TokenKind
	Value interface{}
}

func NewToken(kind TokenKind, value interface{}) *Token {
	return &Token{
		Name:  Codex[kind],
		Kind:  kind,
		Value: value,
	}
}
