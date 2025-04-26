package evaluator

import (
	"Lugo/parser"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestAssignNumber(t *testing.T) {
	code := `
		local a=4
	`
	parser, err := participle.Build[parser.Lua]()
	if err != nil {
		t.Fatal(err)
	}
	tr, err := parser.ParseString("test", code)
	if err != nil {
		t.Fatal(err)
	}
	eval := NewEval(*tr)
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("a")
	if e != nil {
		t.Fatal(e)
	}
	if value.(*Int).value != 4 {
		t.Fatalf("Should have 4 in 'a', instead it has '%v'", value.(*Int).value)
	}
}

func TestSumNumbers(t *testing.T) {
	code := `
		local a=4+3
	`
	parser, err := participle.Build[parser.Lua]()
	if err != nil {
		t.Fatal(err)
	}
	tr, err := parser.ParseString("test", code)
	if err != nil {
		t.Fatal(err)
	}
	eval := NewEval(*tr)
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("a")
	if e != nil {
		t.Fatal(e)
	}
	if value.(*Int).value != 7 {
		t.Fatalf("Should have 7 in 'a', instead it has '%v'", value.(*Int).value)
	}
}

func TestSumAndMultNumbers(t *testing.T) {
	code := `
		local a=4+3*4
	`
	parser, err := participle.Build[parser.Lua]()
	if err != nil {
		t.Fatal(err)
	}
	tr, err := parser.ParseString("test", code)
	if err != nil {
		t.Fatal(err)
	}
	eval := NewEval(*tr)
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("a")
	if e != nil {
		t.Fatal(e)
	}
	if value.(*Int).value != 16 {
		t.Fatalf("Should have 16 in 'a', instead it has '%v'", value.(*Int).value)
	}
}

func TestSumAndMultNumbers2(t *testing.T) {
	code := `
		local a=(4+3)*4
	`
	parser, err := participle.Build[parser.Lua]()
	if err != nil {
		t.Fatal(err)
	}
	tr, err := parser.ParseString("test", code)
	if err != nil {
		t.Fatal(err)
	}
	eval := NewEval(*tr)
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("a")
	if e != nil {
		t.Fatal(e)
	}
	if value.(*Int).value != 28 {
		t.Fatalf("Should have 28 in 'a', instead it has '%v'", value.(*Int).value)
	}
}

func TestDeclareFunc(t *testing.T) {
	code := `
	local a = 10
	function prova(a,b)
		return a+b
	end
	return prova(3,4)
	`
	parser, err := participle.Build[parser.Lua]()
	if err != nil {
		t.Fatal(err)
	}
	tr, err := parser.ParseString("test", code)
	if err != nil {
		t.Fatal(err)
	}
	eval := NewEval(*tr)
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("return")
	if e != nil {
		t.Fatal(e)
	}
	if value.(*Int).value != 7 {
		t.Fatalf("Should have 7 in 'a', instead it has '%v'", value.(*Int).value)
	}
}
