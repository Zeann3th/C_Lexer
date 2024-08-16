package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zeann3th/C_Compiler/internal/core"
)

var (
	lastChar    = byte(32)
	currentByte = 0
)

func main() {
	fileName := "internal/example/hello.cpp"

	source, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(source))
	src := core.NewSource(source)
	for src.Cursor < len(src.Content) {
		fmt.Printf("%v (%v)\n", src.NextToken(), src.Bufnr)
	}
}
