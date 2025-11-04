package parser

type Expression struct {
	KeyWord                  *KeyWordNoExpression `(?!@@)(`
	LambdaFunctionExpression *ExpressionFunction  `@@`
	TableExpression          *TableDeclaration    `|@@`
	MathExpression           *MathExpression      `|@@)`
}

type TableRetrieveWithBracket struct {
	TableName string      `@Ident"["`
	Index     *Expression `@@ "]"` // Parentheses
}

type TableRetrieveWithoutBracket struct {
	TableName *string `@Ident "."`
	Index     *Value  ` @@`
}

type TableDeclaration struct {
	Entries []*TableEntry `"{" @@* "}"` // Parentheses
}

type TableEntry struct {
	Name  *Value      `(@@ "=")?`
	Value *Expression `@@`
	Come  string      `","?`
}

type MathExpression struct {
	HExp *TermExpression `@@`  // Highest level: Terms
	LExp []*LExpression  `@@*` // Lower precedence: Addition & Subtraction
}

type TermExpression struct {
	LeftTerm  *BaseValueExp   `@@`
	Operator  *string         `(@("/" | "*" | ">" "=" | ">" | "<" "=" | "<")` // Multiplication or division
	RightTerm *TermExpression `@@)?`
}

type LExpression struct {
	Operator string          `@("+" | "-" | "or" | "and" | "." "." | "=" "=" | "~" "=" )`
	HExp     *TermExpression `@@`
}

type BaseValueExp struct {
	Base       *Value          `@@`
	Expression *MathExpression `| "(" @@ ")"` // Parentheses
}

type Variable struct {
	Visibility *string `@"local"?`
	Name       string  `@Ident "="`
}

type Value struct {
	Int                         *int                         `@Int`
	Float                       *float32                     `|@Float`
	FunctionCall                *FunctionCall                `|@@`
	String                      *string                      `|@String`
	Bool                        *string                      `|@("true" | "false") `
	TableRetrieveWithoutBracket *TableRetrieveWithoutBracket `|@@`
	TableRetrieveWithBracket    *TableRetrieveWithBracket    `|@@`
	Identifier                  *string                      `|@Ident`
}

type FunctionCall struct {
	Name string               `@Ident`
	Args []*ParamFunctionCall `"("@@*")"`
}

type ParamFunctionCall struct {
	Param *Expression `@@`
	Coma  *string     `","?`
}

type ExpressionFunction struct {
	Declaration string                      `@"function"`
	Args        []*ParamFunctionDeclaration `"("@@*")"`
	Body        Lua                         `@@`
	Return      *ReturnExpression           `@@?`
	End         string                      `"end"!`
}
type ParamFunctionDeclaration struct {
	Param string  `@Ident`
	Coma  *string `","?`
}
