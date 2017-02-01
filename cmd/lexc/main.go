package main

import (
	"fmt"
	"github.com/keizo042/goc/parser"
)

func main() {
	lexer := parser.NewLexer(" 1 + 1 ")
	lexer.Lex()
	for {
		i := lexer.NextItem()
		fmt.Println(i)
		if i.Typ == lex.ItemEOF {

			break
		}
	}
}
