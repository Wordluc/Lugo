package parser

type Statement struct {
	NoStatement       *KeyWordNoStatement `(?!@@)(`
	StatementFunction *StatementFunction  `@@`
	StatementVariable *StatementVariable  `|@@`
	ReturnExpresion   *ReturnExpresion    `|@@)`
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
