package eval

import (
	"Lugo/parser"
	"fmt"
)

func (p *Program) EvalFunctionCall(call *parser.FunctionCall) (Value, error) {
	f, ok := p.function[call.Name]
	if !ok {
		return nil, fmt.Errorf("Function %v not registered", call.Name)
	}
	var e error
	var params []Value = make([]Value, len(call.Args))
	for i := range call.Args {
		params[i], e = p.EvalExp(*call.Args[i].Param)
		if e != nil {
			return nil, e
		}
	}
	return f.Call(params...)
}
