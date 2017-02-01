%{
package parser

import (
	"github.com/keizo042/goc/ast"
)


%}


%union {
    token ast.Item
}

%token<token> DIGIT STRING PLUS MINUS DIV MULTI

%%

%%

type Parser {

}

func  New(lexer *lex.Lexer) *Parser {
}

func (p *Parser)Run()(*ast.Tree, error) {

}
