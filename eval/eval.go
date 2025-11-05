package eval

import (
	"Lugo/parser"
	"errors"
)

type Program struct {
	*Environment
	parser.Lua
}

func NewEval(tree parser.Lua) *Program {
	return NewCustomEval(tree, NewEnvironment())
}
func NewCustomEval(tree parser.Lua, env *Environment) *Program {
	return &Program{
		env,
		tree,
	}
}
func NewHigherTempEval(higher *Program, dict *Dictionary) *Program {
	env := NewEnvironment()
	env.SetHigherEnvironment(higher.Environment)
	for keyRaw, v := range dict.Elements {
		if keyRaw.Type() != StringType {
			continue
		}
		key := keyRaw.(*String)
		if v.Type() == FunctionType {
			env.global[key.value] = v
		} else {
			env.AddVariable(key.value, v)
		}
	}
	return &Program{
		env,
		higher.Lua,
	}
}

func (p *Program) Run() error {
	for _, st := range p.Lua.Statements {
		if e := p.EvalStatement(st); e != nil {
			return e
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
func (p *Program) EvalStatement(st *parser.Statement) error {
	switch {
	case st.StatementVariable != nil:
		v := st.StatementVariable
		value, e := p.EvalExp(v.Expression)
		if e != nil {
			return e
		}
		found := p.SetVariable(v.Variable.Name, value)
		if found {
			return nil
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
	case st.StatementIfCondition != nil:
		var continueIteration bool
		var e error
		if continueIteration, e = p.EvalBodyUnderCondition(&st.StatementIfCondition.Condition, st.StatementIfCondition.Body); e != nil {
			return e
		}
		if !continueIteration {
			return nil
		}
		for _, elseIf := range st.StatementIfCondition.ElseIf {
			if continueIteration, e = p.EvalBodyUnderCondition(elseIf.Condition, *elseIf.Body); e != nil {
				return e
			}
			if !continueIteration {
				return nil
			}
		}
		if st.StatementIfCondition.Else != nil {
			if _, e := p.EvalBodyUnderCondition(nil, *st.StatementIfCondition.Else); e != nil {
				return e
			}
		}
	case st.StatementFor != nil:
		env := NewEnvironment()
		env.higherEnv = p.Environment
		bodyResult := NewCustomEval(st.StatementFor.Body, p.Environment)
		from, e := p.EvalExp(st.StatementFor.From.Expression)
		if e != nil {
			return e
		}
		e = p.AddVariable(st.StatementFor.From.Variable.Name, from)
		if e != nil {
			return e
		}
		to, e := p.EvalValue(&st.StatementFor.To)
		if e != nil {
			return e
		}
		for {
			e = bodyResult.Run()
			if e != nil {
				return e
			}
			fromCondition, e := bodyResult.GetRawVariable(st.StatementFor.From.Variable.Name)
			if e != nil {
				return e
			}
			res, e := fromCondition.EvalOp("==", to)
			if e != nil {
				return e
			}
			if res.(*Bool).value {
				break
			}
			var stepper Value = &Int{value: 1}
			if st.StatementFor.Step != nil {
				stepper, e = bodyResult.EvalValue(st.StatementFor.Step)
				if e != nil {
					return e
				}
			}
			newFrom, e := fromCondition.EvalOp("+", stepper)
			if e != nil {
				return e
			}
			bodyResult.SetVariable(st.StatementFor.From.Variable.Name, newFrom)

		}
	}
	return nil
}
func (p *Program) EvalBodyUnderCondition(condition *parser.Expression, body parser.Lua) (continueIteration bool, err error) {
	var conditionResult Value = &Bool{value: true}
	if condition != nil {
		conditionResult, err = p.EvalExp(*condition)
		if err != nil {
			return false, err
		}
		if conditionResult.Type() != BoolType {
			return false, errors.New("if condition has to be a bool")
		}
	}
	res := conditionResult.(*Bool)

	if res.value {
		env := NewEnvironment()
		env.higherEnv = p.Environment
		bodyResult := NewCustomEval(body, p.Environment)
		return false, bodyResult.Run()
	}
	return true, nil
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
		nil,
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
		nil,
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
