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
	value, e := eval.GetRawVariable("a")
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
	value, e := eval.GetRawVariable("a")
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
	value, e := eval.GetRawVariable("a")
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
	value, e := eval.GetRawVariable("a")
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
	value, e := eval.GetRawVariable("return")
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
		value, _ := eval.GetRawVariable(key)
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
	value, _ := eval.GetRawVariable("a")
	if value.(*String).value != "ciao luca" {
		t.Fatalf("%v should be %v, instead is %v", "a", "ciao luca", value.(*String).value)
	}

	value, _ = eval.GetRawVariable("b")
	if value.(*Bool).value != true {
		t.Fatalf("%v should be %v, instead is %v", "b", "true", value.(*Bool).value)
	}

	value, _ = eval.GetRawVariable("c")
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
	value, _ := eval.GetRawVariable("a")
	if value.(*Int).value != 9 {
		t.Fatalf("%v should be %v, instead is %v", "a", 9, value.(*Int).value)
	}
}
func TestDictionary(t *testing.T) {
	code := `
	local f = {
		a="hello",
		b="world",
		c={
			"all",
		}
	}
	c=f.a .. " " .. f.b .. " " .. f.c[1]
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
	value, e := eval.GetRawVariable("c")
	if e != nil {
		t.Fatal(e)
	}
	if r, _ := value.(*String); r.value != "hello world all" {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be %v, instead is %v", "c", "hello world", r.value)
	}
}
func TestIfCondition(t *testing.T) {
	code := `
	local res = false
	if a==4 then
		res=true
	end
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
	eval.AddVariable("a", &Int{value: 4})
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("res")
	if e != nil {
		t.Fatal(e)
	}
	if r, _ := value.(bool); !r {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be %v, instead is %v", "res", "true", r)
	}
}
func TestIfCondition2(t *testing.T) {
	code := `
	local res = "ciao"
	if a==4 then
		res="hello"
	else
		res="dio"
	end
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
	eval.AddVariable("a", &Int{value: 3})
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("res")
	if e != nil {
		t.Fatal(e)
	}
	if r, _ := value.(string); r != "dio" {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be %v, instead is %v", "res", "dio", r)
	}
}
func TestIfConditionWithElseIf(t *testing.T) {
	code := `
	local res = "ciao"
	if a==4 then
		res="hello"
	elseif a==2 then
		res="dio"
	elseif a==1 then
		res="11111"
	else 
		res="porcoo"
	end
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
	eval.AddVariable("a", &Int{value: 1})
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("res")
	if e != nil {
		t.Fatal(e)
	}
	if r, _ := value.(string); r != "11111" {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be %v, instead is %v", "res", "11111", r)
	}
}
func TestIfConditionWithElseIf2(t *testing.T) {
	code := `
	local res = "ciao"
	if a==4 then
		res="hello"
	elseif a==2 then
		res="dio"
	elseif a==1 then
		res="11111"
	else 
		res="porcoo"
	end
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
	eval.AddVariable("a", &Int{value: 2})
	e := eval.Run()
	if e != nil {
		t.Fatal(e)
	}
	value, e := eval.GetVariable("res")
	if e != nil {
		t.Fatal(e)
	}
	if r, _ := value.(string); r != "dio" {
		if e != nil {
			t.Error(e)
		}
		t.Fatalf("%v should be %v, instead is %v", "res", "dio", r)
	}
}
