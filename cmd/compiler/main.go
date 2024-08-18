package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zeann3th/C_Compiler/internal/lexer"
)

var (
	lastChar    = byte(32)
	currentByte = 0
)

func main() {
	fileName := "internal/example/gcd.c"

	source, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("[FILEIO]: ERROR: Failed to open file <%v>", fileName))
	}
	// fmt.Println(string(source))
	src := lexer.NewLexer(source)
	fmt.Printf("%v\t%v\t\t(%v)\n", "Id", "Name", "Value")
	for src.Cursor < len(src.Content) {
		token := src.NextToken()
		// core.ExpectToken(src, core.TOKEN_SYMBOL)
		fmt.Printf("%v\t%v\t\t(%v)\n", token.Kind, token.Name, token.Value)
	}
}
