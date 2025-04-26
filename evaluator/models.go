package evaluator

import (
	"Lugo/parser"
	"fmt"
)

type TypeValue string

const (
	IntType    TypeValue = "int"
	FloatType  TypeValue = "float"
	StringType TypeValue = "string"
	BoolType   TypeValue = "bool"
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
func (i *Int) EvalOp(op string, v Value) (Value, error) {
	switch v := v.(type) {
	case *Int:
		return EvalInts(i, op, v)
	case *Float:
		return EvalFloats(i.CastFloat(), op, v)
	case *String, *Bool:
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

func (i *Float) EvalOp(op string, v Value) (Value, error) {
	switch v := v.(type) {
	case *Float:
		return EvalFloats(i, op, v)
	case *Int:
		return EvalFloats(i, op, v.CastFloat())
	case *String, *Bool:
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
	case *Float, *Int, *Bool:
		return nil, fmt.Errorf("The operation %v isnt possible with type %v and %v", op, i.Type(), v.Type())
	case *String:
		if op == ".." {
			return &String{
				i.value + v.value,
			}, nil
		}
	}
	return nil, fmt.Errorf("The operation %v isnt defined for type %v and %v", op, i.Type(), v.Type())
}

type Bool struct {
	value bool
}

func (i *Bool) Type() TypeValue {
	return BoolType
}
func (i *Bool) EvalOp(op string, v Value) (Value, error) {
	return nil, fmt.Errorf("The operation %v isnt defined for type %v and %v", op, i.Type(), v.Type())
}
func EvalFloats(a *Float, op string, b *Float) (*Float, error) {
	var value float32
	switch {
	case op == "+":
		value = a.value + b.value
	case op == "-":
		value = a.value - b.value
	case op == "*":
		value = a.value * b.value
	case op == "/":
		value = a.value / b.value
	default:
		return nil, fmt.Errorf("The operation %v isnt defined for type %v and %v", op, a.Type(), b.Type())
	}
	return &Float{value}, nil
}
func EvalInts(a *Int, op string, b *Int) (*Int, error) {
	var value int
	switch {
	case op == "+":
		value = a.value + b.value
	case op == "-":
		value = a.value - b.value
	case op == "*":
		value = a.value * b.value
	case op == "/":
		value = a.value / b.value
	default:
		return nil, fmt.Errorf("The operation %v isnt defined for type %v and %v", op, a.Type(), b.Type())
	}
	return &Int{value}, nil
}

type Function struct {
	Body    parser.Lua
	Params  []string
	BaseEnv *Environment
}

func (f *Function) Call(params ...Value) (Value, error) {
	fun := NewEval(f.Body)
	if e := fun.SetHigherEnvironment(f.BaseEnv); e != nil {
		return nil, e
	}
	for i := range f.Params {
		if i >= len(params) {
			fun.AddVariable(f.Params[i], nil)
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
