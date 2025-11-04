package eval

import (
	"Lugo/parser"
	"errors"
	"fmt"
)

type TypeValue string

const (
	IntType        TypeValue = "int"
	FloatType      TypeValue = "float"
	StringType     TypeValue = "string"
	BoolType       TypeValue = "bool"
	FunctionType   TypeValue = "function"
	DictionaryType TypeValue = "Dictionary"
)

type Value interface {
	Type() TypeValue
	EvalOp(string, Value) (Value, error)
}

type Int struct {
	value int
}

func (i *Int) Type() TypeValue {
	return IntType
}

func (i *Int) CastFloat() *Float {
	return &Float{
		float32(i.value),
	}
}

func (i *Int) CastString() *String {
	return &String{
		fmt.Sprint(i.value),
	}
}

func (i *Int) EvalOp(op string, v Value) (Value, error) {
	switch v := v.(type) {
	case *Int:
		return EvalInts(i, op, v)
	case *Float:
		return EvalFloats(i.CastFloat(), op, v)
	case *String:
		return EvalStrings(i.CastString(), op, v)
	case *Bool:
		return nil, fmt.Errorf("The operation %v isnt possible with type %v and %v", op, i.Type(), v.Type())
	}
	return nil, fmt.Errorf("The operation %v isnt defined for type %v and %v", op, i.Type(), v.Type())
}

type Float struct {
	value float32
}

func (i *Float) Type() TypeValue {
	return FloatType
}

func (i *Float) CastString() *String {
	return &String{
		fmt.Sprint(i.value),
	}
}

func (i *Float) EvalOp(op string, v Value) (Value, error) {
	switch v := v.(type) {
	case *Float:
		return EvalFloats(i, op, v)
	case *Int:
		return EvalFloats(i, op, v.CastFloat())
	case *String:
		return EvalStrings(i.CastString(), op, v)
	case *Bool:
		return nil, fmt.Errorf("The operation %v isnt possible with type %v and %v", op, i.Type(), v.Type())
	}
	return nil, fmt.Errorf("The operation %v isnt defined for type %v and %v", op, i.Type(), v.Type())
}

type String struct {
	value string
}

func (i *String) Type() TypeValue {
	return StringType
}

func (i *String) EvalOp(op string, v Value) (Value, error) {
	switch v := v.(type) {
	case *Int:
		return EvalStrings(i, op, v.CastString())
	case *Float:
		return EvalStrings(i, op, v.CastString())
	case *String:
		return EvalStrings(i, op, v)
	}
	return nil, fmt.Errorf("The operation %v isnt possible with type %v and %v", op, i.Type(), v.Type())
}

type Bool struct {
	value bool
}

func (i *Bool) Type() TypeValue {
	return BoolType
}

func (i *Bool) EvalOp(op string, v Value) (Value, error) {
	if v.Type() != i.Type() {
		return nil, fmt.Errorf("The operation %v isnt possible with type %v and %v", op, i.Type(), v.Type())
	}

	if op == "==" {
		return &Bool{i.value == v.(*Bool).value}, nil
	}
	if op == "~=" {
		return &Bool{i.value != v.(*Bool).value}, nil
	}

	return nil, fmt.Errorf("The operation %v isnt defined for type %v and %v", op, i.Type(), v.Type())
}

func EvalFloats(a *Float, op string, b *Float) (Value, error) {
	switch op {
	case "+":
		return &Float{a.value + b.value}, nil
	case "-":
		return &Float{a.value - b.value}, nil
	case "*":
		return &Float{a.value * b.value}, nil
	case "/":
		return &Float{a.value / b.value}, nil
	}
	switch op {
	case ">=":
		return &Bool{a.value >= b.value}, nil
	case "<=":
		return &Bool{a.value <= b.value}, nil
	case ">":
		return &Bool{a.value > b.value}, nil
	case "<":
		return &Bool{a.value < b.value}, nil
	case "==":
		return &Bool{a.value == b.value}, nil
	case "~=":
		return &Bool{a.value != b.value}, nil
	}
	return nil, fmt.Errorf("The operation %v isn't defined for type %v and %v", op, a.Type(), b.Type())
}

func EvalInts(a *Int, op string, b *Int) (Value, error) {
	switch op {
	case "+":
		return &Int{a.value + b.value}, nil
	case "-":
		return &Int{a.value - b.value}, nil
	case "*":
		return &Int{a.value * b.value}, nil
	case "/":
		return &Float{float32(a.value) / float32(b.value)}, nil
	}
	switch op {
	case ">=":
		println(a.value)
		return &Bool{a.value >= b.value}, nil
	case "<=":
		return &Bool{a.value <= b.value}, nil
	case ">":
		return &Bool{a.value > b.value}, nil
	case "<":
		return &Bool{a.value < b.value}, nil
	case "==":
		return &Bool{a.value == b.value}, nil
	case "~=":
		return &Bool{a.value != b.value}, nil
	}
	return nil, fmt.Errorf("The operation %v isn't defined for type %v and %v", op, a.Type(), b.Type())
}

func EvalStrings(a *String, op string, b *String) (Value, error) {
	//remove " from the begin and the end
	//	a.value = a.value[1 : len(a.value)-1]
	//	b.value = b.value[1 : len(b.value)-1]

	if op == ".." {
		return &String{a.value + b.value}, nil
	}
	switch op {
	case ">=":
		return &Bool{a.value >= b.value}, nil
	case "<=":
		return &Bool{a.value <= b.value}, nil
	case ">":
		return &Bool{a.value > b.value}, nil
	case "<":
		return &Bool{a.value < b.value}, nil
	case "==":
		return &Bool{a.value == b.value}, nil
	case "~=":
		return &Bool{a.value != b.value}, nil
	}
	return nil, fmt.Errorf("The operation %v isn't defined for type %v and %v", op, a.Type(), b.Type())
}

type Function struct {
	Body    parser.Lua
	Params  []string
	BaseEnv *Environment
}

func (f *Function) Type() TypeValue {
	return FunctionType
}

func (f *Function) EvalOp(op string, v Value) (Value, error) {
	return nil, errors.New("Function does't support this operation:" + op)
}

func (f *Function) Call(params ...Value) (Value, error) {
	fun := NewEval(f.Body)
	if e := fun.SetHigherEnvironment(f.BaseEnv); e != nil {
		return nil, e
	}
	for i := range f.Params {
		if i >= len(params) {
			fun.AddVariable(f.Params[i], nil) //to see
		}
		fun.AddVariable(f.Params[i], params[i])
	}
	err := fun.Run()
	if err != nil {
		return nil, err
	}
	value, err := fun.Environment.GetVariable("return")
	if err != nil {
		return nil, err
	}
	return value, err
}

type Dictionary struct {
	//TODO; refactor
	Elements map[Value]Value
}

// TODO; refactor
func (i *Dictionary) Get(key Value) (res Value, e error) {
	for i, v := range i.Elements {
		if i.Type() != key.Type() {
			continue
		}
		if res, e = i.EvalOp("==", key); e == nil {
			if res.(*Bool).value {
				return v, nil
			}

		}
		if e != nil {
			return nil, e
		}
	}
	return res, nil
}
func (i *Dictionary) Type() TypeValue {
	return DictionaryType
}

func (i *Dictionary) EvalOp(op string, v Value) (Value, error) {
	return nil, errors.New("Dictionary does't support this operation:" + op)
}
