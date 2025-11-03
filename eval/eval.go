package eval

import (
	"Lugo/parser"
)

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
			if v.Variable.Visibility != nil && *v.Variable.Visibility == "local" {
				if e := p.Environment.AddVariable(v.Variable.Name, value); e != nil {
					return e
				}
			}
			if e := p.Environment.AddGlobalVariable(v.Variable.Name, value); e != nil {
				return e
			}
		case st.StatementFunction != nil:
			v := st.StatementFunction
			f := p.getFunction(v)
			if e := p.Environment.AddFunction(v.Name, f); e != nil {
				return e
			}
		}
	}
	for _, exp := range p.Lua.Expression {
		_, e := p.EvalExp(*exp)
		if e != nil {
			return e
		}
	}
	if p.ReturnExpression != nil {
		v := p.ReturnExpression.ValueReturned
		value, e := p.EvalExp(*v)
		if e != nil {
			return e
		}
		p.Environment.AddVariable("return", value)
		return nil
	}

	return nil
}
func (p *Program) getFunction(exp *parser.StatementFunction) Function {
	args := make([]string, len(exp.Args))
	for i := range exp.Args {
		args[i] = exp.Args[i].Param
	}
	return Function{
		exp.Body,
		args,
		p.Environment,
	}
}
func (p *Program) getLambdaFunction(exp *parser.ExpressionFunction) *Function {
	args := make([]string, len(exp.Args))
	for i := range exp.Args {
		args[i] = exp.Args[i].Param
	}
	return &Function{
		exp.Body,
		args,
		p.Environment,
	}
}

func (p *Program) EvalExp(exp parser.Expression) (Value, error) {
	if m := exp.MathExpression; m != nil {
		return p.EvalMath(m)
	}
	if m := exp.LambdaFunctionExpression; m != nil {
		return p.getLambdaFunction(m), nil
	}
	if m := exp.TableExpression; m != nil {
		return p.EvalTable(m)
	}

	return nil, nil
}
