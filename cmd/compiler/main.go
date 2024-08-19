package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zeann3th/C_Compiler/internal/lexer"
	"github.com/zeann3th/C_Compiler/internal/parser"
)

var (
	lastChar    = byte(32)
	currentByte = 0
)

func main() {
	fileName := "internal/example/hello.c"

	source, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("[FILEIO]: ERROR: Failed to open file <%v>", fileName))
	}
	// Tokenize(source)

	Parse(source)
}

func Tokenize(src []byte) {
	l := lexer.NewLexer(src)
	fmt.Printf("%v\t%v\t\t(%v)\n", "Id", "Name", "Value")
	for l.Cursor < len(l.Content) {
		token := l.NextToken()
		fmt.Printf("%v\t%v\t\t(%v)\n", token.Kind, token.Name, token.Value)
	}
}

func Parse(src []byte) {
	l := lexer.NewLexer(src)
	p := parser.NewParser(l)
	p.Cursor = 0
	for p.Cursor < len(p.Content) {
		p.GetNextToken()
		p.ParsePrimary()
	}
}
