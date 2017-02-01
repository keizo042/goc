package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func min(l, r int) int {
	if l > r {
		return r
	} else {
		return l
	}
}

func compare(expect, actual []Item) (int, bool) {
	i := 0
	l := len(expect)
	r := len(actual)
	if l != r {
		fmt.Fprintf(os.Stderr, "expect:%d,actual:%d", l, r)
	}
	for {

		lval := expect[i]
		rval := actual[i]
		if lval.Typ == ItemEOF && rval.Typ == ItemEOF {
			return 0, true
		}

		if lval.Typ == ItemEOF || rval.Typ == ItemEOF {
			return i, false
		}

		if lval.Typ != rval.Typ {
			return i, false
		}
		if lval.Token != rval.Token {
			return i, false
		}
		i++
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
	var actual []Item
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ItemEOF {
			break
		}

	}

	expect := []Item{
		Item{Token: "(", Typ: ItemParenL, line: 1},
		Item{Token: "1", Typ: ItemDigit, line: 1},
		Item{Token: "+", Typ: ItemPlus, line: 1},
		Item{Token: "2", Typ: ItemDigit, line: 1},
		Item{Token: ")", Typ: ItemParenR, line: 1},
		Item{Token: "*", Typ: ItemMulti, line: 1},
		Item{Token: "2", Typ: ItemDigit, line: 1},
		Item{Token: "+", Typ: ItemPlus, line: 1},
		Item{Token: "(", Typ: ItemParenL, line: 1},
		Item{Token: "4", Typ: ItemDigit, line: 1},
		Item{Token: "/", Typ: ItemDiv, line: 1},
		Item{Token: "2", Typ: ItemDigit, line: 1},
		Item{Token: ")", Typ: ItemParenR, line: 1},
		Item{Token: "", Typ: ItemEOF, line: 2},
	}
	fmt.Println(len(expect))

	i, b := compare(expect, actual)
	if !b {
		if i < 0 {
			t.Errorf("fail! expected:%v\n", expect)
		} else {
			t.Errorf("fail! error %d\n", i)
			for i, e := range expect {
				t.Errorf("expect:%d\t, %s", i, e)
			}
			t.Error("")
			for i, e := range actual {
				t.Errorf("actual:%d\t, %s", i, e)
			}
		}
	}

}

func TestLex001(t *testing.T) {
	s, err := i(testsrc[1])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	var actual []Item
	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ItemEOF {
			break
		}

	}
	expect := []Item{
		Item{Token: "100", Typ: ItemDigit},
		Item{Token: "", Typ: ItemEOF},
	}
	i, b := compare(expect, actual)
	if !b {
		if i < 0 {
			t.Errorf("fail! expected:%v\n", expect)
		} else {
			t.Errorf("fail! error %d\n", i)
			for i, e := range expect {
				t.Errorf("expect:%d\t, %s", i+1, e)
			}
			t.Error("")
			for i, e := range actual {
				t.Errorf("actual:%d\t, %s", i+1, e)
			}
		}
	}

}

func TestLex002(t *testing.T) {
	s, err := i(testsrc[2])
	if err != nil {
		t.Fatal(err.Error())
	}
	var actual []Item
	lexer := New(s)
	expect := []Item{
		Item{Token: "1", Typ: ItemDigit},
		Item{Token: "+", Typ: ItemPlus},
		Item{Token: "2", Typ: ItemDigit},
		Item{Token: "", Typ: ItemEOF},
	}

	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ItemEOF {
			break
		}

	}
	i, b := compare(expect, actual)
	if !b {
		if i < 0 {
			t.Errorf("fail! expected:%v\n", expect)
		} else {
			t.Errorf("fail! error %d\n", i)
			for i, e := range expect {
				t.Errorf("expect:%d\t, %s", i+1, e)
			}
			t.Error("")
			for i, e := range actual {
				t.Errorf("actual:%d\t, %s", i+1, e)
			}
		}
	}
}

/*
func TestLex003(t *testing.T) {
	s, err := i(testsrc[3])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	expect := []Item{
		Item{Token: "string", Typ: ItemEOF},
		Item{Token: "", Typ: ItemEOF},
	}
	var actual []Item
	lexer.Lex()
	go func() {
		for {
			select {
			case e := <-lexer.Items:
				actual = append(actual, e)
			}

		}
	}()
	<-lexer.Done
	i, b := compare(expect, actual)
	if !b {
		if i < 0 {
			t.Errorf("fail! expected:%v\n", expect)
		} else {
			t.Errorf("fail! error %d\n", i)
			for i, e := range expect {
				t.Errorf("expect:%d\t, %s", i+1, e)
			}
			t.Error("")
			for i, e := range actual {
				t.Errorf("actual:%d\t, %s", i+1, e)
			}
		}
	}
}
*/

func TestLex004(t *testing.T) {
	s, err := i(testsrc[4])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	expect := []Item{
		Item{Token: "6", Typ: ItemDigit},
		Item{Token: "/", Typ: ItemDiv},
		Item{Token: "2", Typ: ItemDigit},
		Item{Token: "", Typ: ItemEOF},
	}
	var actual []Item
	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ItemEOF {
			break
		}

	}
	i, b := compare(expect, actual)
	if !b {
		if i < 0 {
			t.Errorf("fail! expected:%v\n", expect)
		} else {
			t.Errorf("fail! error %d\n", i)
			for i, e := range expect {
				t.Errorf("expect:%d\t, %s", i+1, e)
			}
			t.Error("")
			for i, e := range actual {
				t.Errorf("actual:%d\t, %s", i+1, e)
			}
		}
	}
}

func TestLex005(t *testing.T) {
	s, err := i(testsrc[5])
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	var actual []Item
	expect := []Item{
		Item{Token: "", Typ: ItemEOF},
	}
	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ItemEOF {
			break
		}

	}
	i, b := compare(expect, actual)
	if !b {
		if i < 0 {
			t.Errorf("fail! expected:%v\n", expect)
		} else {
			t.Errorf("fail! error %d\n", i)
			for i, e := range expect {
				t.Errorf("expect:%d\t, %s", i+1, e)
			}
			t.Error("")
			for i, e := range actual {
				t.Errorf("actual:%d\t, %s", i+1, e)
			}
		}
	}
}
