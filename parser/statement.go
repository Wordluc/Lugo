package parser

type Statement struct {
	NoStatement          *KeyWordNoStatement   `(?!@@)(`
	StatementFunction    *StatementFunction    `@@`
	StatementVariable    *StatementVariable    `|@@`
	StatementIfCondition *StatementIfCondition `|@@`
	StatementFor         *StatementFor         `|@@)`
}
type StatementFunction struct {
	Declaration string                      `@"function"`
	Name        string                      `@Ident`
	Args        []*ParamFunctionDeclaration `"("@@*")"`
	Body        Lua                         `@@"end"!`
}
type StatementIfCondition struct {
	Condition Expression        `"if" @@ "then"`
	Body      Lua               `@@`
	ElseIf    []StatementElseIf `@@*`
	Else      *Lua              `("else"@@)?`
	End       string            `"end"`
}
type StatementFor struct {
	For  string            `"for"`
	From StatementVariable `@@","`
	To   Value             `@@`
	Step *Value            `(","@@)?`
	Do   string            `"do"`
	Body Lua               `@@`
	End  string            `"end"`
}
type StatementElseIf struct {
	Condition *Expression `"elseif" @@ "then"`
	Body      *Lua        `@@`
}
type StatementElse struct {
	Else     *string `"else"`
	BodyElse *Lua    `@@`
}
type StatementVariable struct {
	Variable   Variable   `@@`
	Expression Expression `@@`
}
