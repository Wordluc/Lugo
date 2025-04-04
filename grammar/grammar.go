package grammar

type ToString interface {
	toString() string
}

type Lua struct {
	Statements []*Statement `@@*` //local a=n;
	Expresions []*Expresion `@@*` //4+2
}

// Define keyword that hasnt be valuated has expresions
type KeyWordNoValue struct {
	KeyWord *string `"end"`
}
type KeyWord struct {
	KeyWord *string `@"end"`
}

func (e *Lua) toString() string {
	res := ""
	for i, ex := range e.Statements {
		res += ex.toString()
		if i != len(e.Statements)-1 {
			res += "\n"
		}
	}
	res += "\n"
	for i, ex := range e.Expresions {
		res += ex.toString()
		if i != len(e.Expresions)-1 {
			res += "\n"
		}
	}
	return res
}
