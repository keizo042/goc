%{
package parser

import (
	"github.com/keizo042/goc/ast"
)


%}

%token<token> num

%union {
    token ast.Item
}


%%

%%

type Parser {

}

func  New(lexer *lex.Lexer) *Parser{
}

func (p *Parser)Run()(*ast.Tree, error) {

}
