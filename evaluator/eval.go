package evaluator

import "Lugo/parser"

type Environment struct {
	Variables map[string]Value
}

func NewEnvironment() Environment {
	return Environment{
		Variables: make(map[string]Value),
	}
}
func (e *Environment) AddVariable(name string, v Value) error {
	e.Variables[name] = v
	return nil
}

func (e *Environment) GetVariable(name string) (Value, error) {
	return e.Variables[name], nil
}

type Program struct {
	Environment
	parser.Lua
}

func NewEval(tree parser.Lua) Program {
	return Program{
		NewEnvironment(),
		tree,
	}
}
func (p *Program) Run() error {
	for _, st := range p.Lua.Statements {
		if v := st.StatementVariable; v != nil {
			value, e := p.EvalExp(v.Expression)
			if e != nil {
				return e
			}
			if e := p.Environment.AddVariable(v.Variable.Name, value); e != nil {
				return e
			}
		}
	}
	return nil
}

func (p *Program) EvalExp(exp parser.Expression) (Value, error) {
	return nil, nil
}
