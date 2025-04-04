package grammar

import "fmt"

type ToString interface {
	toString() string
}

type Lua struct {
	Statements []*Statement `@@*` //local a=n;
	Expresions []*Expresion `@@*` //4+2
}
type Statement struct {
	StatementVariable *StatementVariable `@@`
	StatementFunction *StatementFunction `|@@`
}
type StatementFunction struct {
	Declaration string   `@"function"`
	Name        string   `@Ident`
	Args        []string `"("@Ident*")"`
	Body        Lua      `@@"end"`
}
type StatementVariable struct {
	Variable  Variable  `@@`
	Expresion Expresion `@@`
}

type Variable struct {
	Visibility *string `@"local"?`
	Name       string  `@Ident "="`
}

type Expresion struct {
	HExp *HExpresion   `@@`  // Highest level: Terms
	LExp []*LExpresion `@@*` // Lower precedence: Addition & Subtraction
}

type LExpresion struct {
	Operator string      `@("+" | "-")`
	HExp     *HExpresion `@@`
}

type HExpresion struct {
	BaseValue *BaseValueExp `@@`
	Right     *OpFactor     `@@*` // Lower precedence: Multiplication & Division
}

type OpFactor struct {
	Operator  string        `@("/"|"*")` // Multiplication or division
	BaseValue *BaseValueExp `@@`
}

type BaseValueExp struct {
	Base       *Value     `@@`
	Expression *Expresion `| "(" @@ ")"` // Parentheses
}

type Value struct {
	Number       *float32      ` @Float | @Int`
	FunctionCall *FunctionCall `| @@`
	String       *string       `| @String`
	Bool         *bool         `| @("true" | "false") `
}

type FunctionCall struct {
	Name string          `@Ident`
	Args []*ParmFunction `"("@@*")"`
}

type ParmFunction struct {
	Parm *Expresion `@@`
	Coma *string    `","?`
}

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
func (e *Statement) toString() string {
	if e.StatementVariable != nil {
		return e.StatementVariable.toString()
	}
	if e.StatementFunction != nil {
		return e.StatementFunction.toString()
	}
	return "<undefined>"
}
func (e *Lua) toString() string {
	res := ""
	for i, ex := range e.Statements {
		res += ex.toString()
		if i != len(e.Statements)-1 {
			res += "\n"
		}
	}
	res += "\n"
	for i, ex := range e.Expresions {
		res += ex.toString()
		if i != len(e.Expresions)-1 {
			res += "\n"
		}
	}
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

func (e *StatementVariable) toString() string {
	res := ""
	res += e.Variable.toString() + "="
	res += e.Expresion.toString()
	return res
}

func (e *ParmFunction) toString() string {
	return fmt.Sprint(e.Parm.toString(), ",")
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
