package lex

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func compare(l, r []Item) (int, bool) {
	for i := 0; ; i++ {
		lval := l[i]
		rval := r[i]
		if lval.Typ != rval.Typ {
			return i, false
		}
		if lval.Token != rval.Token {
			return i, false
		}
		if lval.Typ == ItemEOF || rval.Typ == ItemEOF {
			break
		}
	}
	return 0, true
}

func i(src string) (string, error) {
	fp, err := os.Open(src)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", err
	}
	return string(b), nil

}

func consume(t *testing.T, s string, l *Lexer) []Item {
	var elems []Item
	for {
		select {
		case e := <-l.Items:
			elems = append(elems, e)
		case <-time.After(1 * time.Second):
			t.Fatal("too late processing")
			return nil
		}

	}
	return elems
}

var testsrc = []string{
	"./testdata/test000.calc",
	"./testdata/test001.calc",
	"./testdata/test002.calc",
	"./testdata/test003.calc",
	"./testdata/test004.calc",
	"./testdata/test005.calc",
}

func TestLex000(t *testing.T) {
	s, err := i(testsrc[0])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	var actual []Item
	go func() {
		actual = consume(t, testsrc[0], lexer)
	}()
	<-lexer.Done

	expect := []Item{
		Item{Token: "(", Typ: ItemParenL},
		Item{Token: "1", Typ: ItemDigit},
		Item{Token: "+", Typ: ItemPlus},
		Item{Token: "2", Typ: ItemDigit},
		Item{Token: ")", Typ: ItemParenR},
		Item{Token: "*", Typ: ItemMulti},
		Item{Token: "2", Typ: ItemDigit},
		Item{Token: "+", Typ: ItemPlus},
		Item{Token: "(", Typ: ItemParenL},
		Item{Token: "4", Typ: ItemDigit},
		Item{Token: "/", Typ: ItemDiv},
		Item{Token: "2", Typ: ItemDigit},
		Item{Token: ")", Typ: ItemParenR},
		Item{Token: "", Typ: ItemEOF},
	}

	i, b := compare(actual, expect)
	if !b {
		t.Errorf("fail! expected:%v, actual:%v", expect[i], actual[i])
	}

}

func TestLex001(t *testing.T) {
	s, err := i(testsrc[1])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	actual := consume(t, testsrc[1], lexer)
	expect := []Item{
		Item{Token: "", Typ: ItemEOF},
	}
	n, b := compare(expect, actual)
	if b {
		t.Error(n)
	}

}

func TestLex002(t *testing.T) {
	s, err := i(testsrc[2])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	actual := consume(t, testsrc[2], lexer)
	expect := []Item{
		Item{Token: "", Typ: ItemEOF},
	}
	n, b := compare(expect, actual)
	if b {
		t.Error(n)
	}
}

func TestLex003(t *testing.T) {
	s, err := i(testsrc[3])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	actual := consume(t, testsrc[3], lexer)
	expect := []Item{
		Item{Token: "", Typ: ItemEOF},
	}
	n, b := compare(expect, actual)
	if b {
		t.Error(n)
	}
}

func TestLex004(t *testing.T) {
	s, err := i(testsrc[4])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	actual := consume(t, testsrc[4], lexer)
	expect := []Item{
		Item{Token: "", Typ: ItemEOF},
	}
	n, b := compare(expect, actual)
	if b {
		t.Error(n)
	}
}

func TestLex005(t *testing.T) {
	s, err := i(testsrc[5])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	actual := consume(t, testsrc[5], lexer)
	expect := []Item{
		Item{Token: "", Typ: ItemEOF},
	}
	n, b := compare(expect, actual)
	if b {
		t.Error(n)
	}
}
