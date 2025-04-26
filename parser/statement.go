package parser

type Statement struct {
	NoStatement       *KeyWordNoStatement `(?!@@)(`
	StatementFunction *StatementFunction  `@@`
	StatementVariable *StatementVariable  `|@@)`
}
type StatementFunction struct {
	Declaration string                      `@"function"`
	Name        string                      `@Ident`
	Args        []*ParamFunctionDeclaration `"("@@*")"`
	Body        Lua                         `@@"end"!`
}

type StatementVariable struct {
	Variable   Variable   `@@`
	Expression Expression `@@`
}
