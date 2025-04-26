package eval

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
func TestLogicalOperation(t *testing.T) {
	code := `
	local a = 10>4
	b=4==4
	c=6<3
	d=4<=4
	e=0>=4
	f=3>=4==false
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
	if e != nil {
		t.Fatal(e)
	}
	var result = map[string]bool{
		"a": true,
		"b": true,
		"c": false,
		"d": true,
		"e": false,
		"f": true,
	}
	for key, v := range result {
		value, _ := eval.GetVariable(key)
		if value == nil {
			t.Fatalf("%v not found", key)
		}
		if value.(*Bool).value != v {
			t.Fatalf("%v should be %v, instead is %v", key, v, value.(*Bool).value)
		}
	}
}
func TestStringOperation(t *testing.T) {
	code := `
	local a = "ciao " .. "luca" 
	local b = "xbc" > "a"
	local c = "abc" == "abc"
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
	if e != nil {
		t.Fatal(e)
	}
	value, _ := eval.GetVariable("a")
	if value.(*String).value != "ciao luca" {
		t.Fatalf("%v should be %v, instead is %v", "a", "ciao luca", value.(*String).value)
	}

	value, _ = eval.GetVariable("b")
	if value.(*Bool).value != true {
		t.Fatalf("%v should be %v, instead is %v", "b", "true", value.(*Bool).value)
	}

	value, _ = eval.GetVariable("c")
	if value.(*Bool).value != true {
		t.Fatalf("%v should be %v, instead is %v", "c", "true", value.(*Bool).value)
	}
}
