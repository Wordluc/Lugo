package parser

type ToString interface {
	toString() string
}

type Lua struct {
	Statements       []*Statement      `@@*` //local a=n;
	Expression       []*Expression     `@@*` //4+2
	ReturnExpression *ReturnExpression `@@?`
}

// Define keyword that hasnt be valuated has expression
type KeyWordNoExpression struct {
	KeyWord *string `"end"|"return"`
}

// Define keyword that hasnt be valuated has statement
type KeyWordNoStatement struct {
	KeyWord *string `"return"`
}

type ReturnExpression struct {
	ValueReturned *Expression `"return" @@`
}

func (e *Lua) toString() string {
	res := ""
	for i, ex := range e.Statements {
		res += ex.toString()
		if i != len(e.Statements)-1 {
			res += "\n"
		}
		if len(e.Statements)-1 == i {
			res += "\n"
		}
	}
	for i, ex := range e.Expression {
		res += ex.toString()
		if i != len(e.Expression)-1 {
			res += "\n"
		}
	}
	if e.ReturnExpression != nil {
		res += e.ReturnExpression.toString() + "\n"
	}
	return res
}
func (e *ReturnExpression) toString() string {
	res := ""
	res += "return "
	res += e.ValueReturned.toString()
	return res
}
