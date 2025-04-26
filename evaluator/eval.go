package evaluator

import "Lugo/parser"

type Program struct {
	*Environment
	parser.Lua
}

func NewEval(tree parser.Lua) *Program {
	return &Program{
		NewEnvironment(),
		tree,
	}
}

func (p *Program) Run() error {
	for _, st := range p.Lua.Statements {
		switch {
		case st.StatementVariable != nil:
			v := st.StatementVariable
			value, e := p.EvalExp(v.Expression)
			if e != nil {
				return e
			}
			if e := p.Environment.AddVariable(v.Variable.Name, value); e != nil {
				return e
			}
		case st.StatementFunction != nil:
			v := st.StatementFunction
			f := Function{
				v.Body,
				v.Args,
				p.Environment,
			}
			if e := p.Environment.AddFunction(v.Name, f); e != nil {
				return e
			}
		case st.ReturnExpression != nil:
			v := st.StatementVariable
			value, e := p.EvalExp(v.Expression)
			if e != nil {
				return e
			}
			p.Environment.AddVariable("return", value)
			return nil
		}

	}
	return nil
}

func (p *Program) EvalExp(exp parser.Expression) (Value, error) {
	if m := exp.MathExpression; m != nil {
		return p.EvalMath(m)
	}
	return nil, nil
}
