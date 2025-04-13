package parser

type ToString interface {
	toString() string
}

type Lua struct {
	Statements      []*Statement     `@@*` //local a=n;
	Expresions      []*Expresion     `@@*` //4+2
	ReturnExpresion *ReturnExpresion `@@?`
}

// Define keyword that hasnt be valuated has expresions
type KeyWordNoExpresion struct {
	KeyWord *string `"end"|"return"`
}

// Define keyword that hasnt be valuated has statement
type KeyWordNoStatement struct {
	KeyWord *string `"return"`
}

type ReturnExpresion struct {
	ValueReturned *Expresion `"return" @@`
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
	if e.ReturnExpresion != nil {
		res += e.ReturnExpresion.toString() + "\n"
	}
	return res
}
func (e *ReturnExpresion) toString() string {
	res := ""
	res += "return "
	res += e.ValueReturned.toString()
	return res
}
