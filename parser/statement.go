package parser

type Statement struct {
	NoStatement          *KeyWordNoStatement   `(?!@@)(`
	StatementFunction    *StatementFunction    `@@`
	StatementVariable    *StatementVariable    `|@@`
	StatementIfCondition *StatementIfCondition `|@@)`
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
	Else      *StatementElse    `@@?`
	End       string            `"end"`
}
type StatementElseIf struct {
	ElseIf     *Expression `"elseif" @@ "then"`
	BodyElseIf *Lua        `@@`
}
type StatementElse struct {
	Else     *string `"else"`
	BodyElse *Lua    `@@`
}
type StatementVariable struct {
	Variable   Variable   `@@`
	Expression Expression `@@`
}
