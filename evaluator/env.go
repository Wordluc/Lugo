package evaluator

type Environment struct {
	Variables map[string]Value
	Function  map[string]Function
	higherEnv *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		Variables: make(map[string]Value),
	}
}
func (e *Environment) AddVariable(name string, v Value) error {
	e.Variables[name] = v
	return nil
}

func (e *Environment) AddFunction(name string, v Function) error {
	e.Function[name] = v
	return nil
}

func (e *Environment) GetVariable(name string) (Value, error) {
	if v := e.Variables[name]; v != nil {
		return v, nil
	}
	if e.higherEnv == nil {
		return nil, nil
	}
	if v := e.higherEnv.Variables[name]; v != nil {
		return v, nil
	}
	return nil, nil
}

func (e *Environment) SetHigherEnvironment(env *Environment) error {
	e.higherEnv = env
	return nil
}
