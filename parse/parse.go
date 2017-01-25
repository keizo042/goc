package parse

import (
	"github.com/keizo042/goc/lex"
)

type state int

const (
	stateInit state = iota + 1
)

type State struct {
	lexer *lex.Lexer
	src   []lex.Item
	state state
}

func New(l *lex.Lexer) *State {
}

func (s *State) accept(r rune) bool {
}
