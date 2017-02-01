package ast

import (
	"fmt"
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
	Line  uint64
	Pos   int64
	Err   error
}

func (i Item) String() string {
	return fmt.Sprintf("{type:%s\t,token:\"%s\"\t,pos:%d\t,line:%d}", i.Typ, i.Token, i.Pos, i.Line)
}
