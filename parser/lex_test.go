package parser

import (
	"fmt"
	"github.com/keizo042/goc/ast"
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

func compare(expect, actual []ast.Item) (int, bool) {
	i := 0
	l := len(expect)
	r := len(actual)
	if l != r {
		fmt.Fprintf(os.Stderr, "expect:%d,actual:%d", l, r)
	}
	for {

		lval := expect[i]
		rval := actual[i]
		if lval.Typ == ast.ItemEOF && rval.Typ == ast.ItemEOF {
			return 0, true
		}

		if lval.Typ == ast.ItemEOF || rval.Typ == ast.ItemEOF {
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
	var actual []ast.Item
	if err != nil {
		t.Fatal(err.Error())
	}
	lexer := New(s)
	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ast.ItemEOF {
			break
		}

	}

	expect := []ast.Item{
		ast.Item{Token: "(", Typ: ast.ItemParenL, Line: 1},
		ast.Item{Token: "1", Typ: ast.ItemDigit, Line: 1},
		ast.Item{Token: "+", Typ: ast.ItemPlus, Line: 1},
		ast.Item{Token: "2", Typ: ast.ItemDigit, Line: 1},
		ast.Item{Token: ")", Typ: ast.ItemParenR, Line: 1},
		ast.Item{Token: "*", Typ: ast.ItemMulti, Line: 1},
		ast.Item{Token: "2", Typ: ast.ItemDigit, Line: 1},
		ast.Item{Token: "+", Typ: ast.ItemPlus, Line: 1},
		ast.Item{Token: "(", Typ: ast.ItemParenL, Line: 1},
		ast.Item{Token: "4", Typ: ast.ItemDigit, Line: 1},
		ast.Item{Token: "/", Typ: ast.ItemDiv, Line: 1},
		ast.Item{Token: "2", Typ: ast.ItemDigit, Line: 1},
		ast.Item{Token: ")", Typ: ast.ItemParenR, Line: 1},
		ast.Item{Token: "", Typ: ast.ItemEOF, Line: 2},
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
	var actual []ast.Item
	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ast.ItemEOF {
			break
		}

	}
	expect := []ast.Item{
		ast.Item{Token: "100", Typ: ast.ItemDigit},
		ast.Item{Token: "", Typ: ast.ItemEOF},
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
	var actual []ast.Item
	lexer := New(s)
	expect := []ast.Item{
		ast.Item{Token: "1", Typ: ast.ItemDigit},
		ast.Item{Token: "+", Typ: ast.ItemPlus},
		ast.Item{Token: "2", Typ: ast.ItemDigit},
		ast.Item{Token: "", Typ: ast.ItemEOF},
	}

	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ast.ItemEOF {
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
	expect := []ast.Item{
		ast.Item{Token: "string", Typ: ast.ItemEOF},
		ast.Item{Token: "", Typ: ast.ItemEOF},
	}
	var actual []ast.Item
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
	expect := []ast.Item{
		ast.Item{Token: "6", Typ: ast.ItemDigit},
		ast.Item{Token: "/", Typ: ast.ItemDiv},
		ast.Item{Token: "2", Typ: ast.ItemDigit},
		ast.Item{Token: "", Typ: ast.ItemEOF},
	}
	var actual []ast.Item
	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ast.ItemEOF {
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
	var actual []ast.Item
	expect := []ast.Item{
		ast.Item{Token: "", Typ: ast.ItemEOF},
	}
	lexer.Lex()
	for {
		e := lexer.NextItem()
		actual = append(actual, e)
		if e.Typ == ast.ItemEOF {
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
