package eval

import "errors"

type Environment struct {
	local     map[string]Value
	global    map[string]Value
	higherEnv *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		local:  make(map[string]Value),
		global: make(map[string]Value),
	}
}
func (e *Environment) SetVariable(name string, v Value) (found bool) {
	if _, ok := e.local[name]; ok {
		e.local[name] = v
		return true
	}
	if _, ok := e.global[name]; ok {
		e.global[name] = v
		return true
	}
	return false
}
func (e *Environment) AddVariable(name string, v Value) error {
	e.local[name] = v
	return nil
}
func (e *Environment) AddGlobalVariable(name string, v Value) error {
	e.global[name] = v
	return nil
}

func (e *Environment) AddFunction(name string, v Function) error {
	e.global[name] = &v
	return nil
}
func (e *Environment) AddCustomFunction(name string, f func(env *Environment, args []Value) Value) error {
	e.global[name] = &Function{
		customFunction: f,
	}

	return nil
}

func (e *Environment) GetRawVariable(name string) (Value, error) {
	if v := e.local[name]; v != nil {
		return v, nil
	}
	if e.higherEnv != nil {
		if v := e.higherEnv.local[name]; v != nil {
			return v, nil
		}
	}
	if v := e.global[name]; v != nil {
		return v, nil
	}
	return nil, errors.New("Variable " + name + " not found")
}

func (env *Environment) GetVariable(name string) (any, error) {
	value, e := env.GetRawVariable(name)
	if e != nil {
		return nil, e
	}
	switch v := value.(type) {
	case *String:
		return v.Get(), nil
	case *Int:
		return v.Get(), nil
	case *Float:
		return v.Get(), nil
	case *Bool:
		return v.Get(), nil
	}
	return nil, errors.New("Type not supported for getVariable")
}
func (e *Environment) SetHigherEnvironment(base *Environment) error {
	e.higherEnv = base
	e.global = base.global
	return nil
}
