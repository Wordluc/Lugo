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
	Variable  Variable  `@@`
	Expresion Expresion `@@`
}

type Variable struct {
	Visibility *string `@"local"?`
	Name       string  `@Ident "="`
}

type Expresion struct {
	Left  *Value         `@@`
	Right []*ExpresionOp `@@?`
}

type ExpresionOp struct {
	Operation string     `@("+" | "-" | "/")`
	Expresion *Expresion `@@`
}

type Value struct {
	Number       *float32      ` @Float | @Int`
	FunctionCall *FunctionCall `| @@`
	String       *string       `| @String`
	Bool         *bool         `| @("true" | "false") `
}

type FunctionCall struct {
	Name string          `@Ident"("`
	Args []*ParmFunction `@@*")"`
}

type ParmFunction struct {
	Parm *Expresion `@@`
	Coma *string    `","?`
}

func (e *Lua) toString() string {
	res := ""
	for i, ex := range e.Statements {
		res += ex.toString()
		if i != len(e.Statements)-1 {
			res += "\n"
		}
	}
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

func (e *Statement) toString() string {
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

func (e *ExpresionOp) toString() string {
	res := ""
	res += e.Operation
	res += e.Expresion.toString()
	return res
}

func (e *Expresion) toString() string {
	res := ""
	res += e.Left.toString()
	for _, ex := range e.Right {
		res += ex.toString()
	}
	return res
}
