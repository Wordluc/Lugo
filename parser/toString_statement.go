package parser

import "fmt"

func (e *StatementFunction) toString() string {
	res := ""
	res += e.Declaration
	res += " "
	res += e.Name
	res += "("
	for _, v := range e.Args {
		res += v
	}
	res += "){"
	res += "\n"
	res += e.Body.toString()
	res += "\n}"
	return res
}

func (e *ParamFunction) toString() string {
	return fmt.Sprint(e.Param.toString(), ",")
}

func (e *Statement) toString() string {
	if e.StatementVariable != nil {
		return e.StatementVariable.toString()
	}
	if e.StatementFunction != nil {
		return e.StatementFunction.toString()
	}
	return "<undefined>"
}

func (e *StatementVariable) toString() string {
	res := ""
	res += e.Variable.toString() + "="
	res += e.Expression.toString()
	return res
}
