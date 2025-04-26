package parser

import "fmt"

func (e *Expression) toString() string {
	if e.LambdaFunctionExpression != nil {
		return e.LambdaFunctionExpression.toString()
	}
	if e.TableExpression != nil {
		return e.TableExpression.toString()
	}
	if e.MathExpression != nil {
		return e.MathExpression.toString()
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
	if e.TableRetrieveWithoutBracket != nil {
		return e.TableRetrieveWithoutBracket.toString()
	}
	if e.TableRetrieveWithBracket != nil {
		return e.TableRetrieveWithBracket.toString()
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
	if e.Int != nil {
		return fmt.Sprint(*e.Int)
	}
	if e.Float != nil {
		return fmt.Sprint(*e.Float)
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

func (e *MathExpression) toString() string {
	res := e.HExp.toString()
	for _, op := range e.LExp {
		res += op.toString()
	}
	return res
}
func (e *LExpression) toString() string {
	res := ""
	res += " " + e.Operator + " "
	res += e.HExp.toString()
	return res
}
func (t *TermExpression) toString() string {
	res := "("
	if t.LeftTerm != nil {
		res += t.LeftTerm.toString()
	}
	if t.Operator != nil {
		fmt.Printf("jfklsdf %+v \n", *t.Operator)
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
func (e *TableRetrieveWithBracket) toString() string {
	res := ""
	res += e.TableName
	res += "["
	res += e.Index.toString()
	res += "]"
	return res
}
func (e *TableRetrieveWithoutBracket) toString() string {
	res := ""
	res += *e.TableName
	res += "."
	res += *e.Index
	return res
}

func (e *ExpressionFunction) toString() string {
	res := ""
	res += e.Declaration
	res += " "
	res += "("
	for _, v := range e.Args {
		res += v.toString()
	}
	res += "){"
	res += ""
	res += e.Body.toString()
	res += "}"
	return res
}

func (e *ParamFunctionDeclaration) toString() string {
	return fmt.Sprint(e.Param, ",")
}
