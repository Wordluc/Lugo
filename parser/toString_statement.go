package parser

import "fmt"

func (e *StatementFunction) toString() string {
	res := ""
	res += e.Declaration
	res += " "
	res += e.Name
	res += "("
	for _, v := range e.Args {
		res += v.toString()
	}
	res += "){"
	res += "\n"
	res += e.Body.toString()
	res += "\n}"
	return res
}

func (e *StatementIfCondition) toString() string {
	res := "if "
	res += e.Condition.toString()
	res += "then\n"
	res += e.Body.toString()
	res += "\n"
	for i := range e.ElseIf {
		res += "elseif "
		res += e.ElseIf[i].ElseIf.toString()
		res += " then\n"
		res += e.ElseIf[i].BodyElseIf.toString()
		res += "\n"
	}
	if e.Else != nil {
		res += "else\n"
		res += e.Else.BodyElse.toString()
		res += "\n"
	}
	res += "end"
	return res
}
func (e *ParamFunctionCall) toString() string {
	return fmt.Sprint(e.Param.toString(), ",")
}

func (e *Statement) toString() string {
	if e.StatementVariable != nil {
		return e.StatementVariable.toString()
	}
	if e.StatementFunction != nil {
		return e.StatementFunction.toString()
	}
	if e.StatementIfCondition != nil {
		return e.StatementIfCondition.toString()
	}
	return "<undefined>"
}

func (e *StatementVariable) toString() string {
	res := ""
	res += e.Variable.toString() + "="
	res += e.Expression.toString()
	return res
}
