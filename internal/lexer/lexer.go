// Package lexer provides core functions of a C compiler
package lexer

import (
	"strconv"
)

// Source represents the inputed source code from a file

type Lexer struct {
	Content  []byte  // Detailed content of the source code
	Cursor   int     // The current read position, updates after consuming a byte
	LastByte byte    // The last byte that was consumed
	Bufnr    string  // A string buffer if encounter a specific event like keywords, string literals, char literals
	NumBuf   float64 // A number buffer for decimals and floats
	Line     int
	Col      int
}

// NewLexer creates a new instance of type Source
func NewLexer(content []byte) *Lexer {
	return &Lexer{
		Content:  content,
		Cursor:   0,
		LastByte: byte(32),
		Bufnr:    "",
		NumBuf:   0,
		Line:     0,
		Col:      0,
	}
}

func (l *Lexer) PeekByte() byte {
	b := l.Content[l.Cursor]
	return b
}

// GetByte returns the next byte and increment the cursor position.
//
// No saving => in order to save, must implement LastByte = GetByte()
func (l *Lexer) GetByte() byte {
	b := l.Content[l.Cursor]
	l.Cursor++
	l.Col++
	if b == '\n' {
		l.Line++
		l.Col = 0
	}
	return b
}

// GetByteAndSave returns the next byte and increment the cursor position.
//
// Saving the consumed byte to LastByte
func (l *Lexer) GetByteAndSave() byte {
	b := l.GetByte()
	l.LastByte = b
	return b
}

// NextToken returns a number, it may be the code for tokens specified in token.go or the ASCII code for undefined character.
//
// Available tokens: TOKEN_OPERATION, TOKEN_COMMENT, TOKEN_SYMBOL, TOKEN_OPAREN, TOKEN_CPAREN, TOKEN_OCURLY, TOKEN_CCURLY, TOKEN_SEMICOLON, TOKEN_NUMBER, TOKEN_STRING, TOKEN_CHAR, TOKEN_RETURN, TOKEN_EOF, TOKEN_ERROR.
func (l *Lexer) NextToken() *Token {
	tokenHandlers := map[byte]func() *Token{
		'#':  l.HandleOperation,
		';':  l.HandlePunctuation,
		'.':  l.HandlePunctuation,
		',':  l.HandlePunctuation,
		'/':  l.HandleComment,
		'"':  l.HandleStringLiteral,
		'\'': l.HandleCharLiteral,
	}

	l.Bufnr = ""
	// Skip spaces and f*ckin windows \r\n
	for IsSpace(l.LastByte) || l.LastByte == '\n' || l.LastByte == '\r' {
		if l.Cursor < len(l.Content) {
			l.LastByte = l.GetByte()
		} else {
			return &Token{
				Name: Codex[EOF],
				Kind: EOF,
			}
		}
	}

	if handler, ok := tokenHandlers[l.LastByte]; ok {
		return handler()
	}
	if IsOperand(l.LastByte) {
		return l.HandleLogicalExpr()
	}
	if IsAlpha(l.LastByte) {
		return l.HandleKeyword()
	}
	if IsDigit(l.LastByte) || l.LastByte == '.' {
		return l.HandleNumber()
	}
	if IsOC(l.LastByte) {
		return l.HandleBracket()
	}
	return l.HandleUnknownCharacter()
}

// HandleOperation checks for '#' and consumes all bytes into Bufnr until a <CR>.
func (l *Lexer) HandleOperation() *Token {
	l.Bufnr = ""
	for l.LastByte != 0 && l.LastByte != '\r' {
		l.Bufnr += string(l.LastByte)
		l.GetByteAndSave()
	}
	l.GetByteAndSave()
	return NewToken(OPERATION, l.Bufnr)
}

// HandlePunctuation checks for ';' and consumes it into Bufnr.
//
// Zeann3th: Maybe it should skip this and check for semicolon in AST and parser step...
func (l *Lexer) HandlePunctuation() *Token {
	tmp := map[string]TokenKind{
		";": SEMICOLON,
		",": COMMA,
		".": DOT,
	}
	l.Bufnr = string(l.LastByte)
	l.GetByteAndSave()
	return NewToken(tmp[l.Bufnr], l.Bufnr)
}

// HandleComment checks for inline/block comments.
//
// For inline comments: It checks '/' 2 times and consumes all bytes until end of line.
//
// For block comments: It checks for '/' followed by '*' and consumes all bytes until encountering '*' and '/' consecutively.
func (l *Lexer) HandleComment() *Token {
	l.GetByteAndSave()
	if l.LastByte == '/' {
		for l.LastByte != '\n' && l.LastByte != 0 {
			l.GetByteAndSave()
		}
		return NewToken(COMMENT, l.Bufnr)
	} else if l.LastByte == '*' {
		for {
			l.GetByteAndSave()
			if l.LastByte == '*' {
				l.GetByteAndSave()
				if l.LastByte == '/' {
					l.GetByteAndSave()
					return NewToken(COMMENT, l.Bufnr)
				}
			}
		}
	}
	return NewToken(ERROR, nil)
}

// HandleKeyword checks for specific keywords. For example: return, extern, continue, break,...
// It checks for an alphabetic character and consumes all bytes until encountering a non-alphabetic character like '\n' or numbers.
//
// Zeann3th: At the moment, i only implement "return", not yet finished... :(
func (l *Lexer) HandleKeyword() *Token {
	l.Bufnr = string(l.LastByte)
	for IsAlNum(l.GetByteAndSave()) {
		l.Bufnr += string(l.LastByte)
	}
	kindHandlers := map[string]TokenKind{
		"return": RETURN,
		"int":    TYPE,
		"void":   TYPE,
		"char":   TYPE,
		"string": TYPE,
		"float":  TYPE,
		"if":     KEYWORD,
		"for":    KEYWORD,
	}

	if kind, ok := kindHandlers[l.Bufnr]; ok {
		return NewToken(kind, l.Bufnr)
	}

	return NewToken(SYMBOL, l.Bufnr)
}

// HandleStringLiteral checks for '"' and consumes all bytes until another '"'. Just a normal string procedure.
func (l *Lexer) HandleStringLiteral() *Token {
	l.Bufnr = ""
	for l.Cursor < len(l.Content) {
		l.GetByteAndSave()
		if l.LastByte == '"' {
			l.GetByteAndSave()
			return NewToken(STRING, l.Bufnr)
		}
		l.Bufnr += string(l.LastByte)
	}
	return NewToken(ERROR, nil)
}

// HandleCharLiteral checks for ”' and consumses a byte.
func (l *Lexer) HandleCharLiteral() *Token {
	l.GetByteAndSave()
	l.Bufnr = string(l.LastByte)
	if l.GetByteAndSave() == '\'' {
		return NewToken(CHAR, l.Bufnr)
	}
	return NewToken(ERROR, nil)
}

// HandleNumber checks for a number or a '.' for floats.
func (l *Lexer) HandleNumber() *Token {
	numStr := ""
	for do := true; do; do = (IsDigit(l.LastByte) || l.LastByte == '.') {
		numStr += string(l.LastByte)
		l.LastByte = l.GetByte()
	}
	numTmp, _ := strconv.ParseFloat(numStr, 64)
	l.NumBuf = numTmp
	l.Bufnr = strconv.FormatFloat(numTmp, 'f', -1, 64)
	return NewToken(NUMBER, l.Bufnr)
}

// HandleUnknownCharacter checks for character that does not satisfy any scenario.
func (l *Lexer) HandleUnknownCharacter() *Token {
	current := l.LastByte
	l.Bufnr = string(current)
	l.LastByte = l.GetByte()
	return NewToken(UNDEFINED, l.Bufnr)
}

func (l *Lexer) HandleBracket() *Token {
	tmp := map[string]TokenKind{
		"(": OPAREN,
		")": CPAREN,
		"{": OCURLY,
		"}": CCURLY,
	}
	l.Bufnr = string(l.LastByte)
	l.LastByte = l.GetByte()
	return NewToken(tmp[l.Bufnr], l.Bufnr)
}

func (l *Lexer) HandleLogicalExpr() *Token {
	var kind TokenKind
	l.Bufnr = string(l.LastByte)
	switch l.LastByte {
	case '+':
		kind = ADD
	case '-':
		kind = SUB
	case '/':
		kind = DIV
	case '*':
		kind = MUL
	case '=':
		if tmp := l.PeekByte(); tmp == '=' {
			l.Cursor++
			l.LastByte = tmp
			l.Bufnr += string(l.LastByte)
			kind = ASSIGN
		} else {
			kind = EQ
		}
	case '<':
		if tmp := l.PeekByte(); tmp == '=' {
			l.Cursor++
			l.LastByte = tmp
			l.Bufnr += string(l.LastByte)
			kind = LTE
		} else if tmp == '>' {
			l.Cursor++
			l.LastByte = tmp
			l.Bufnr += string(l.LastByte)
			kind = NOTEQ
		} else {
			kind = LT
		}
	case '>':
		if tmp := l.PeekByte(); tmp == '=' {
			l.Cursor++
			l.LastByte = tmp
			l.Bufnr += string(l.LastByte)
			kind = GTE
		} else {
			kind = GT
		}
	case '!':
		if tmp := l.PeekByte(); tmp == '=' {
			l.Cursor++
			l.LastByte = tmp
			l.Bufnr += string(l.LastByte)
			kind = NOTEQ
		} else {
			kind = NOT
		}
	}
	l.GetByteAndSave()
	return NewToken(kind, l.Bufnr)
}
