package core

import (
	"strconv"
)

type Source struct {
	Content  []byte
	Cursor   int
	LastByte byte
	Bufnr    string
	NumBuf   float64
}

func NewSource(content []byte) *Source {
	return &Source{
		Content:  content,
		Cursor:   0,
		LastByte: byte(32),
		Bufnr:    "",
		NumBuf:   0,
	}
}

func (s *Source) GetByte() byte {
	b := s.Content[s.Cursor]
	s.Cursor++
	return b
}

func (s *Source) GetByteAndSave() byte {
	b := s.GetByte()
	s.LastByte = b
	return b
}

func (s *Source) NextToken() int {
	for IsSpace(s.LastByte) || s.LastByte == 3 || s.LastByte == '\n' {
		s.LastByte = s.GetByte()
	}

	// #include, #define
	if s.LastByte == '#' {
		s.Bufnr = string(s.LastByte)
		s.GetByteAndSave()
		return TOKEN_HASH
	}

	// Semicolon
	if s.LastByte == ';' {
		s.Bufnr = string(s.LastByte)
		s.GetByteAndSave()
		return TOKEN_SEMICOLON
	}

	// Comment
	if s.LastByte == '/' {
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
	}

	// Keyswords like int, return, string
	if IsAlpha(s.LastByte) {
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

	// String literals
	if s.LastByte == '"' {
		s.Bufnr = ""
		for s.Cursor < len(s.Content) {
			s.GetByteAndSave()
			if s.LastByte == '"' {
				s.GetByteAndSave()
				return TOKEN_STRING
			}
			s.Bufnr += string(s.LastByte)
		}
	}

	// Char literals
	if s.LastByte == '\'' {
		s.GetByteAndSave()
		s.Bufnr = string(s.LastByte)
		if s.GetByteAndSave() == '\'' {
			return TOKEN_CHAR
		}
		return TOKEN_ERROR
	}

	// Nums, Decimals
	if IsDigit(s.LastByte) || s.LastByte == '.' {
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

	// Parens and Curlies (){}
	if ISOC(s.LastByte) {
		s.Bufnr = string(s.LastByte)
		s.GetByteAndSave()
		if s.Bufnr == "(" {
			return TOKEN_OPAREN
		}
		if s.Bufnr == ")" {
			return TOKEN_CPAREN
		}
		if s.Bufnr == "{" {
			return TOKEN_OCURLY
		}
		if s.Bufnr == "}" {
			return TOKEN_CCURLY
		}
	}

	// EOF
	if s.LastByte == 0 {
		return TOKEN_EOF
	}

	// Return unknown byte
	current := s.LastByte
	s.Bufnr = string(current)
	s.LastByte = s.GetByte()
	return int(current)
}

func IsSpace(b byte) bool {
	return b == 32
}

func IsAlpha(b byte) bool {
	return (b >= 97 && b <= 122) || (b >= 65 && b <= 90)
}

func IsDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func IsAlNum(b byte) bool {
	return IsAlpha(b) || IsDigit(b)
}

func ISOC(b byte) bool {
	return b == 40 || b == 41 || b == 123 || b == 125
}
