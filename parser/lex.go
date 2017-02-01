package parser

import (
	"github.com/keizo042/goc/ast"
	"unicode"
	"unicode/utf8"
)

type stateFn func(*Lexer) stateFn

// Lexer is a Lexical anaysis state Machine. it is recommended to run as gorutine.
type Lexer struct {
	// Items is reciever
	Items chan ast.Item

	state stateFn
	src   string
	start uint64
	pos   uint64
	width uint64

	size uint64
	line uint64
}

func New(src string) *Lexer {
	return &Lexer{
		Items: make(chan ast.Item, 2),

		state: lexText,
		src:   src,
		start: 0,
		pos:   0,
		width: 0,

		size: uint64(len(src)),
		line: 1,
	}
}

func (l *Lexer) NextItem() ast.Item {
	i := <-l.Items
	return i
}

func (l *Lexer) Drain() {
	for range l.Items {
	}
}

func (l *Lexer) Lex() {
	go l.lex()
}

func (l *Lexer) lex() {

	for l.state != nil {
		l.state = l.state(l)
	}

}

func (l *Lexer) eof() bool {
	return l.pos >= l.size
}

func (l *Lexer) next() rune {
	r, s := utf8.DecodeRuneInString(l.src[l.pos:])
	l.width = uint64(s)
	l.pos += l.width
	return r
}

func (l *Lexer) backup() {
	l.pos -= l.width
}

func (l *Lexer) emit(typ ast.ItemType) {
	if typ == ast.ItemEOF {
		l.Items <- ast.Item{Typ: typ}
		return
	}
	token := l.src[l.start:l.pos]
	l.start = l.pos
	l.Items <- ast.Item{
		Token: token,
		Typ:   typ,
		Line:  l.line,
		Pos:   int64(l.pos % l.line),
	}

}

func lexText(l *Lexer) stateFn {
	for {
		c := l.next()
		if l.eof() {
			l.emit(ast.ItemEOF)
			return nil
		}
		switch c {
		case ' ', '\t':
			l.start = l.pos
		case '\n':
			l.start = l.pos
			l.line++
		case '+':
			l.emit(ast.ItemPlus)
		case '-':
			l.emit(ast.ItemMinus)
		case '*':
			l.emit(ast.ItemMulti)
		case '/':
			l.emit(ast.ItemDiv)
		case '(':
			l.emit(ast.ItemParenL)
		case ')':
			l.emit(ast.ItemParenR)
		default:
			l.backup()
			return lexIdent

		}
	}
}

func lexIdent(l *Lexer) stateFn {
	for {
		c := l.next()
		if unicode.IsDigit(c) {
			return lexDigit
		} else {
			return lexToken
		}
	}
}

func lexToken(l *Lexer) stateFn {
	for {
		c := l.next()
		if unicode.IsDigit(c) && 'a' <= c && c <= 'z' && 'A' <= c && c <= 'Z' {
			continue
		}
		l.backup()
		l.emit(ast.ItemIdent)
		break
	}
	return lexText
}

func lexDigit(l *Lexer) stateFn {
	for {
		c := l.next()
		if !unicode.IsDigit(c) {
			l.backup()
			break
		}

	}
	l.emit(ast.ItemDigit)
	return lexText
}
