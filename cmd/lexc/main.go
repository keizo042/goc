package main

import (
	"fmt"
	"github.com/keizo042/goc/lex"
)

func main() {
	lexer := lex.New(" 1 + 1 ")
	lexer.Lex()
	for {
		i := lexer.NextItem()
		fmt.Println(i)
		if i.Typ == lex.ItemEOF {

			break
		}
	}
}
