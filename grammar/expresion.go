package grammar

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
