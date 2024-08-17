// Package core provides core functions of a C compiler
package core

import (
	"strconv"
)

// Source represents the inputed source code from a file

type Source struct {
	Content  []byte  // Detailed content of the source code
	Cursor   int     // The current read position, updates after consuming a byte
	LastByte byte    // The last byte that was consumed
	Bufnr    string  // A string buffer if encounter a specific event like keywords, string literals, char literals
	NumBuf   float64 // A number buffer for decimals and floats
}

// NewSource creates a new instance of type Source
func NewSource(content []byte) *Source {
	return &Source{
		Content:  content,
		Cursor:   0,
		LastByte: byte(32),
		Bufnr:    "",
		NumBuf:   0,
	}
}

// GetByte returns the next byte and increment the cursor position.
//
// No saving => in order to save, must implement LastByte = GetByte()
func (s *Source) GetByte() byte {
	b := s.Content[s.Cursor]
	s.Cursor++
	return b
}

// GetByteAndSave returns the next byte and increment the cursor position.
//
// Saving the consumed byte to LastByte
func (s *Source) GetByteAndSave() byte {
	b := s.GetByte()
	s.LastByte = b
	return b
}

// NextToken returns a number, it may be the code for tokens specified in token.go or the ASCII code for undefined character.
//
// Available tokens: TOKEN_OPERATION, TOKEN_COMMENT, TOKEN_SYMBOL, TOKEN_OPAREN, TOKEN_CPAREN, TOKEN_OCURLY, TOKEN_CCURLY, TOKEN_SEMICOLON, TOKEN_NUMBER, TOKEN_STRING, TOKEN_CHAR, TOKEN_RETURN, TOKEN_EOF, TOKEN_ERROR.
func (s *Source) NextToken() int {
	s.Bufnr = ""
	// Skip spaces and f*ckin windows \r\n
	for IsSpace(s.LastByte) || s.LastByte == '\n' || s.LastByte == '\r' {
		if s.Cursor < len(s.Content) {
			s.LastByte = s.GetByte()
		} else {
			return TOKEN_EOF
		}
	}
	if s.LastByte == '#' {
		return s.HandleOperation()
	}
	if s.LastByte == ';' {
		return s.HandleSemicolon()
	}
	if s.LastByte == '/' {
		return s.HandleComment()
	}
	if IsAlpha(s.LastByte) {
		return s.HandleKeyword()
	}
	if s.LastByte == '"' {
		return s.HandleStringLiteral()
	}
	if s.LastByte == '\'' {
		return s.HandleCharLiteral()
	}
	if IsDigit(s.LastByte) || s.LastByte == '.' {
		return s.HandleNumber()
	}
	return s.HandleUnknownCharacter()
}

// HandleOperation checks for '#' and consumes all bytes into Bufnr until a <CR>.
func (s *Source) HandleOperation() int {
	s.Bufnr = ""
	for s.LastByte != 0 && s.LastByte != '\r' {
		s.Bufnr += string(s.LastByte)
		s.GetByteAndSave()
	}
	s.GetByteAndSave()
	return TOKEN_OPERATION
}

// HandleSemicolon checks for ';' and consumes it into Bufnr.
//
// Zeann3th: Maybe it should skip this and check for semicolon in AST and parser step...
func (s *Source) HandleSemicolon() int {
	s.Bufnr = string(s.LastByte)
	s.GetByteAndSave()
	return TOKEN_SEMICOLON
}

// HandleComment checks for inline/block comments.
//
// For inline comments: It checks '/' 2 times and consumes all bytes until end of line.
//
// For block comments: It checks for '/' followed by '*' and consumes all bytes until encountering '*' and '/' consecutively.
func (s *Source) HandleComment() int {
	s.GetByteAndSave()
	if s.LastByte == '/' {
		for s.LastByte != '\n' && s.LastByte != 0 {
			s.GetByteAndSave()
		}
		return TOKEN_COMMENT
	} else if s.LastByte == '*' {
		for {
			s.GetByteAndSave()
			if s.LastByte == '*' {
				s.GetByteAndSave()
				if s.LastByte == '/' {
					s.GetByteAndSave()
					return TOKEN_COMMENT
				}
			}
		}
	}
	return TOKEN_ERROR
}

// HandleKeyword checks for specific keywords. For example: return, extern, continue, break,...
// It checks for an alphabetic character and consumes all bytes until encountering a non-alphabetic character like '\n' or numbers.
//
// Zeann3th: At the moment, i only implement "return", not yet finished... :(
func (s *Source) HandleKeyword() int {
	s.Bufnr = string(s.LastByte)
	for IsAlNum(s.GetByteAndSave()) {
		s.Bufnr += string(s.LastByte)
	}

	if s.Bufnr == "return" {
		return TOKEN_RETURN
	}
	if s.Bufnr == "fn" {
	}
	return TOKEN_SYMBOL
}

// HandleStringLiteral checks for '"' and consumes all bytes until another '"'. Just a normal string procedure.
func (s *Source) HandleStringLiteral() int {
	s.Bufnr = ""
	for s.Cursor < len(s.Content) {
		s.GetByteAndSave()
		if s.LastByte == '"' {
			s.GetByteAndSave()
			return TOKEN_STRING
		}
		s.Bufnr += string(s.LastByte)
	}
	return TOKEN_ERROR
}

// HandleCharLiteral checks for ”' and consumses a byte.
func (s *Source) HandleCharLiteral() int {
	s.GetByteAndSave()
	s.Bufnr = string(s.LastByte)
	if s.GetByteAndSave() == '\'' {
		return TOKEN_CHAR
	}
	return TOKEN_ERROR
}

// HandleNumber checks for a number or a '.' for floats.
func (s *Source) HandleNumber() int {
	numStr := ""
	for do := true; do; do = (IsDigit(s.LastByte) || s.LastByte == '.') {
		numStr += string(s.LastByte)
		s.LastByte = s.GetByte()
	}
	numTmp, _ := strconv.ParseFloat(numStr, 64)
	s.NumBuf = numTmp
	s.Bufnr = strconv.FormatFloat(numTmp, 'f', -1, 64)
	return TOKEN_NUMBER
}

// HandleUnknownCharacter checks for character that does not satisfy any scenario.
func (s *Source) HandleUnknownCharacter() int {
	current := s.LastByte
	s.Bufnr = string(current)
	s.LastByte = s.GetByte()
	return int(current)
}
