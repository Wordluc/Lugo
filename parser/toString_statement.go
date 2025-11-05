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

func (e *StatementFor) toString() string {
	res := "FOR "
	res += e.From.toString()
	res += ","
	res += e.To.toString()
	if e.Step != nil {
		res += ","
		res += e.Step.toString()
	}
	res += " DO\n"
	res += e.Body.toString()
	res += "\nEND"
	return res
}
func (e *StatementIfCondition) toString() string {
	res := "IF "
	res += e.Condition.toString()
	res += "THEN\n"
	res += e.Body.toString()
	res += "\n"
	for i := range e.ElseIf {
		res += "ELSEIF "
		res += e.ElseIf[i].Condition.toString()
		res += " THEN\n"
		res += e.ElseIf[i].Body.toString()
		res += "\n"
	}
	if e.Else != nil {
		res += "ELSE\n"
		res += e.Else.toString()
		res += "\n"
	}
	res += "END"
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
	if e.StatementFor != nil {
		return e.StatementFor.toString()
	}
	return "<undefined>"
}

func (e *StatementVariable) toString() string {
	res := ""
	res += e.Variable.toString() + "="
	res += e.Expression.toString()
	return res
}
