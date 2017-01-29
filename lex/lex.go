package lex

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type ItemType int64

func (i ItemType) String() string {
	switch i {
	case ItemEOF:
		return "EOF"
	case ItemDigit:
		return "Digit"
	case ItemPlus:
		return "Plus"
	case ItemMinus:
		return "Minus"
	case ItemDiv:
		return "Div"
	case ItemMulti:
		return "Multi"
	case ItemParenL:
		return "ParenL"
	case ItemParenR:
		return "ParenR"
	case ItemIdent:
		return "Ident"
	default:
		return fmt.Sprintf("unknown:%d", i)
	}
}

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

func (i Item) String() string {
	return fmt.Sprintf("{type:%s\t,token:\"%s\"\t,pos:%d\t,line:%d}", i.Typ, i.Token, i.pos, i.line)
}

// Lexer is a Lexical anaysis state Machine. it is recommended to run as gorutine.
type Lexer struct {
	// Items is reciever
	Items chan Item

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
		Items: make(chan Item, 2),

		state: lexText,
		src:   src,
		start: 0,
		pos:   0,
		width: 0,

		size: uint64(len(src)),
		line: 1,
	}
}

func (l *Lexer) NextItem() Item {
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

func (l *Lexer) emit(typ ItemType) {
	if typ == ItemEOF {
		l.Items <- Item{Typ: typ}
		return
	}
	token := l.src[l.start:l.pos]
	l.start = l.pos
	l.Items <- Item{
		Token: token,
		Typ:   typ,
		line:  l.line,
		pos:   int64(l.pos % l.line),
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
			l.start = l.pos
		case '\n':
			l.start = l.pos
			l.line++
		case '+':
			l.emit(ItemPlus)
		case '-':
			l.emit(ItemMinus)
		case '*':
			l.emit(ItemMulti)
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
	l.emit(ItemDigit)
	return lexText
}
