package parser

import "fmt"

func (e *Expresion) toString() string {
	if e.LambdaFunctionExpresion != nil {
		return e.LambdaFunctionExpresion.toString()
	}
	if e.TableExpresion != nil {
		return e.TableExpresion.toString()
	}
	if e.MathExpresion != nil {
		return e.MathExpresion.toString()
	}
	return "<undefined>"
}

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
	if e.TableRetriveWithoutBracket != nil {
		return e.TableRetriveWithoutBracket.toString()
	}
	if e.TableRetriveWithBracket != nil {
		return e.TableRetriveWithBracket.toString()
	}
	if e.Identifier != nil {
		return *e.Identifier
	}
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
	res := ""
	if e.Base != nil {
		res += e.Base.toString()
	}
	if e.Expression != nil {
		res += "(" + e.Expression.toString() + ")"
	}
	return res
}

func (e *MathExpresion) toString() string {
	res := e.HExp.toString()
	for _, op := range e.LExp {
		res += op.toString()
	}
	return res
}
func (e *LExpresion) toString() string {
	res := ""
	res += " " + e.Operator + " "
	res += e.HExp.toString()
	return res
}
func (t *TermExpresion) toString() string {
	res := "("
	if t.LeftTerm != nil {
		res += t.LeftTerm.toString()
	}
	if t.Operator != nil {
		res += " " + *t.Operator
	}
	if t.RightTerm != nil {
		res += " " + t.RightTerm.toString()
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

func (e *TableDeclaration) toString() string {
	res := ""
	res += "{"
	for _, v := range e.Entries {
		res += v.toString() + ","
	}
	res += "}"
	return res
}

func (e *TableEntry) toString() string {
	res := ""
	if e.Name != nil {
		res += *e.Name
		res += "="
	}
	res += e.Value.toString()
	return res
}
func (e *TableRetriveWithBracket) toString() string {
	res := ""
	res += e.TableName
	res += "["
	res += e.Index.toString()
	res += "]"
	return res
}
func (e *TableRetriveWithoutBracket) toString() string {
	res := ""
	res += *e.TableName
	res += "."
	res += *e.Index
	return res
}

func (e *ExpresionFunction) toString() string {
	res := ""
	res += e.Declaration
	res += " "
	res += "("
	for _, v := range e.Args {
		res += v
	}
	res += "){"
	res += ""
	res += e.Body.toString()
	res += "}"
	return res
}
