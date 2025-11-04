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
	g=3>=4~=false
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
		"g": false,
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
	if value.(*String).value != "\"ciao luca\"" {
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
func TestLambdafunction(t *testing.T) {
	code := `
	local f = function () 
		return 8
	end
	a=f()+1
	a+4
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
	if value.(*Int).value != 9 {
		t.Fatalf("%v should be %v, instead is %v", "a", 9, value.(*Int).value)
	}

}
func TestDictionary(t *testing.T) {
	code := `
	local f = {
		a=function ()return 45 end,
		"ciao",
		d=4,
		43,
	}
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
	value, e := eval.GetVariable("f")
	if e != nil {
		t.Fatal(e)
	}
	dic := value.(*Dictionary)
	if r, e := dic.Get(&Int{value: 1}); r.Type() != StringType || e != nil {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be type %v, instead is %v", "1", "string", value.Type())
	}
	if r, e := dic.Get(&Int{value: 2}); r.Type() != IntType || e != nil {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be type %v, instead is %v", "2", "int", value.Type())
	}
	if r, e := dic.Get(&String{value: "a"}); r.Type() != FunctionType || e != nil {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be type %v, instead is %v", "a", "function", value.Type())
	}
	if r, e := dic.Get(&String{value: "d"}); r.Type() != IntType || e != nil {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be type %v, instead is %v", "d", "int", value.Type())
	}

}
