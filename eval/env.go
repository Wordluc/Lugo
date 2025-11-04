package eval

import "errors"

type Environment struct {
	variables map[string]Value
	global    map[string]Value
	higherEnv *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		variables: make(map[string]Value),
		global:    make(map[string]Value),
	}
}
func (e *Environment) AddVariable(name string, v Value) error {
	e.variables[name] = v
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
	if v := e.variables[name]; v != nil {
		return v, nil
	}
	if e.higherEnv != nil {
		if v := e.higherEnv.variables[name]; v != nil {
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
		return v.value, nil
	case *Int:
		return v.value, nil
	case *Float:
		return v.value, nil
	case *Bool:
		return v.value, nil
	}
	return nil, errors.New("Type not supported for getVariable")
}
func (e *Environment) SetHigherEnvironment(base *Environment) error {
	e.higherEnv = base
	e.global = base.global
	return nil
}
