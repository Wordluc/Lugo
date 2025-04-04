package grammar

type Statement struct {
	StatementFunction *StatementFunction `@@`
	StatementVariable *StatementVariable `|@@`
}
type StatementFunction struct {
	Declaration string   `@"function"`
	Name        string   `@Ident`
	Args        []string `"("@Ident*")"`
	Body        Lua      `@@"end"!`
}

type StatementVariable struct {
	Variable  Variable  `@@`
	Expresion Expresion `@@`
}
