package parser

type Expression struct {
	KeyWord                  *KeyWordNoExpression `(?!@@)(`
	LambdaFunctionExpression *ExpressionFunction  `@@`
	TableExpression          *TableDeclaration    `|@@`
	MathExpression           *MathExpression      `|@@)`
}

type TableRetrieve struct {
	TableName       string           `@Ident`
	IndexValue      *TableValueIndex `("."@@`
	IndexExpression *Expression      `|"["@@"]")`
}
type TableValueIndex struct {
	FunctionCall  *FunctionCall  `@@`
	TableRetrieve *TableRetrieve `|@@`
	Identifier    *string        `|@Ident`
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
	IntNegative   *int           `"-"@Int`
	Int           *int           `|@Int`
	Float         *float32       `|@Float`
	String        *string        `|@String`
	Bool          *string        `|@("true" | "false") `
	FunctionCall  *FunctionCall  `|@@`
	TableRetrieve *TableRetrieve `|@@`
	Identifier    *string        `|@Ident`
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
