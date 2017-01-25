package lex

import (
	"unicode"
	"unicode/utf8"
)

type ItemType int64

type stateFn func(*Lexer) stateFn

const (
	ItemEOF ItemType = iota + 1
	ItemDigit
	ItemPlus
	ItemMinus
	ItemDiv
	ItemMulti
	ItemParenL
	ItemParenR
	ItemIdent
)

type Item struct {
	Typ   ItemType
	Token string
	line  uint64
	pos   int64
	err   error
}

// Lexer is a Lexical anaysis state Machine. it is recommended to run as gorutine.
type Lexer struct {
	// Items is reciever
	Items chan Item
	Done  chan bool

	state stateFn
	src   string
	start uint64
	pos   uint64
	width uint64
	chr   rune

	size uint64
	line uint64
}

func New(src string) *Lexer {
	return &Lexer{
		Items: make(chan Item, 2),
		Done:  make(chan bool, 1),

		state: lexText,
		src:   src,
		start: 0,
		pos:   0,
		width: 0,

		size: uint64(len(src)),
		line: 0,
	}
}

func (l *Lexer) Lex() {
	l.lex()
}

func (l *Lexer) lex() {

	for l.state != nil {
		l.state = l.state(l)
	}

}

func (l *Lexer) eof() bool {
	return l.start >= l.size
}

func (l *Lexer) next() rune {
	r, s := utf8.DecodeRuneInString(l.src[l.start+l.pos:])
	l.width = uint64(s)
	l.chr = r
	l.pos += uint64(s)
	return r
}

func (l *Lexer) peek() rune {
	return l.chr
}

func (l *Lexer) backup() rune {
	if l.pos <= 0 {
		l.start -= l.width
	} else {
		l.pos -= l.width
	}
	r, s := utf8.DecodeRuneInString(l.src[l.start+l.pos:])

	l.width = uint64(s)
	l.chr = r
	return r
}

func (l *Lexer) emit(typ ItemType) {
	if typ == ItemEOF {
		l.Items <- Item{Typ: typ}
		l.Done <- true
		return
	}
	token := l.src[l.start : l.start+l.pos]
	l.start += l.pos
	l.pos = 0
	l.Items <- Item{
		Token: token,
		Typ:   typ,
		line:  l.line,
		pos:   int64(l.start % l.line),
	}

}

func lexText(l *Lexer) stateFn {
	for {
		c := l.next()
		if l.eof() {
			l.emit(ItemEOF)
			return nil
		}
		switch c {
		case ' ', '\t':
		case '\n':
			l.line++
		case '+':
			l.emit(ItemPlus)
		case '-':
			l.emit(ItemMinus)
		case '*':
			l.emit(ItemMinus)
		case '/':
			l.emit(ItemDiv)
		case '(':
			l.emit(ItemParenL)
		case ')':
			l.emit(ItemParenR)
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
		l.emit(ItemIdent)
		break
	}
	return nil
}

func lexDigit(l *Lexer) stateFn {
	for {
		c := l.next()
		if !unicode.IsDigit(c) {
			break
		}

	}
	l.emit(ItemDigit)
	return lexText
}
