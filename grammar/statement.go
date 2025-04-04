package grammar

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
