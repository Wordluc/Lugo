package parser

type Expresion struct {
	KeyWord                 *KeyWordNoExpresion `(?!@@)(`
	LambdaFunctionExpresion *ExpresionFunction  `@@`
	TableExpresion          *TableDeclaration   `|@@`
	MathExpresion           *MathExpresion      `|@@)`
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
	HExp *TermExpresion `@@`  // Highest level: Terms
	LExp []*LExpresion  `@@*` // Lower precedence: Addition & Subtraction
}

type TermExpresion struct {
	LeftTerm  *BaseValueExp  `@@`
	Operator  *string        `(@("/"|"*")` // Multiplication or division
	RightTerm *TermExpresion `@@)?`
}

type LExpresion struct {
	Operator string         `@("+" | "-" | "or" | "and")`
	HExp     *TermExpresion `@@`
}

type BaseValueExp struct {
	Base       *Value         `@@`
	Expression *MathExpresion `| "(" @@ ")"` // Parentheses
}

type Variable struct {
	Visibility *string `@"local"?`
	Name       string  `@Ident "="`
}

type Value struct {
	Number                     *float32                    `@Float | @Int`
	FunctionCall               *FunctionCall               `| @@`
	String                     *string                     `| @String`
	Bool                       *bool                       `| @("true" | "false") `
	TableRetriveWithoutBracket *TableRetriveWithoutBracket `|@@`
	TableRetriveWithBracket    *TableRetriveWithBracket    `|@@`
	Identifier                 *string                     `|@Ident`
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
