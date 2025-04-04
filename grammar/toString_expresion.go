package grammar

import "fmt"

func (e *FunctionCall) toString() string {
	res := ""
	res += e.Name + "("
	for _, v := range e.Args {
		res += v.toString()
	}
	res += ")"
	return res
}

func (e *Value) toString() string {
	if e.FunctionCall != nil {
		return e.FunctionCall.toString()
	}
	if e.String != nil {
		return *e.String
	}
	if e.Number != nil {
		return fmt.Sprint(*e.Number)
	}
	if e.Bool != nil {
		return fmt.Sprint(*e.Bool)
	}
	return "<undefined>"
}

func (e *BaseValueExp) toString() string {
	if e.Base != nil {
		return e.Base.toString()
	}
	if e.Expression != nil {
		return e.Expression.toString()
	}
	return ""
}

func (e *Expresion) toString() string {
	res := e.HExp.toString()
	for _, op := range e.LExp {
		res += " " + op.Operator + " " + op.HExp.toString()
	}
	return res
}

func (t *HExpresion) toString() string {
	res := "(" + t.BaseValue.toString()
	if t.Right != nil {
		res += " " + t.Right.Operator + " " + t.Right.BaseValue.toString()
	}
	res += ")"
	return res
}

func (e *Variable) toString() string {
	res := ""
	if e.Visibility != nil {
		res += *e.Visibility + " "
	}
	res += e.Name
	return res
}
