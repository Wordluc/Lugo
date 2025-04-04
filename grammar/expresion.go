package grammar

type Expresion struct {
	KeyWord                    *KeyWordNoExpresion         `(?!@@)(`
	LambdaFunctionExpresion    *ExpresionFunction          `@@`
	TableExpresion             *TableDeclaration           `|@@`
	TableRetriveWithBracket    *TableRetriveWithBracket    `|@@`
	TableRetriveWithoutBracket *TableRetriveWithoutBracket `|@@`
	MathExpresion              *MathExpresion              `|@@)`
}

type TableRetriveWithBracket struct {
	TableName string     `@Ident"["`
	Index     *Expresion `@@ "]"` // Parentheses
}

type TableRetriveWithoutBracket struct {
	TableName *string `@Ident "."`
	Index     *string ` @Ident`
}

type TableDeclaration struct {
	Entries []*TableEntry `"{" @@* "}"` // Parentheses
}

type TableEntry struct {
	Name  *string    `(@Ident "=")?`
	Value *Expresion `@@`
	Come  string     `","?`
}
type MathExpresion struct {
	HExp *HExpresion   `@@`  // Highest level: Terms
	LExp []*LExpresion `@@*` // Lower precedence: Addition & Subtraction
}

type HExpresion struct {
	BaseValue *BaseValueExp `@@`
	Right     *OpFactor     `@@*` // Lower precedence: Multiplication & Division
}

type LExpresion struct {
	Operator string      `@("+" | "-" | "or" "and")`
	HExp     *HExpresion `@@`
}

type OpFactor struct {
	Operator  string        `@("/"|"*")` // Multiplication or division
	BaseValue *BaseValueExp `@@`
}

type BaseValueExp struct {
	Base       *Value     `@@`
	Expression *Expresion `| "(" @@ ")"` // Parentheses
}

type Variable struct {
	Visibility *string `@"local"?`
	Name       string  `@Ident "="`
}

type Value struct {
	Number       *float32      `@Float | @Int`
	FunctionCall *FunctionCall `| @@`
	String       *string       `| @String`
	Bool         *bool         `| @("true" | "false") `
	Identifier   *string       `|@Ident`
}

type FunctionCall struct {
	Name string          `@Ident`
	Args []*ParmFunction `"("@@*")"`
}

type ParmFunction struct {
	Parm *Expresion `@@`
	Coma *string    `","?`
}

type ExpresionFunction struct {
	Declaration string           `@"function"`
	Args        []string         `"("@Ident*")"`
	Body        Lua              `@@`
	Return      *ReturnExpresion `@@?`
	End         string           `"end"!`
}
