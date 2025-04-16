package evaluator

import "Lugo/parser"

func (p *Program) EvalMath(exp *parser.MathExpression) (Value, error) {
	return nil, nil
}
func (p *Program) EvalTempMath(exp *parser.TermExpression) (Value, error) {
	var left Value
	var op string
	var right Value
	var e error
	if exp.LeftTerm.Expression != nil {
		left, e = p.EvalMath(exp.LeftTerm.Expression)
	}
	if exp.LeftTerm.Base != nil {
		left, e = p.EvalValue(exp.LeftTerm.Base)
	}
	if exp.Operator == nil {
		return left, e
	}
	op = *exp.Operator
	right, e = p.EvalTempMath(exp.RightTerm)
	return left.EvalOp(op, right)
}

func (p *Program) EvalValue(exp *parser.Value) (Value, error) { //TableRetrieveWithoutBracket,TableRetrieveWithBracket miss
	switch {
	case exp.Int != nil:
		return &Int{
			value: *exp.Int,
		}, nil
	case exp.Float != nil:
		return &Float{
			value: *exp.Float,
		}, nil
	case exp.String != nil:
		return &String{
			value: *exp.String,
		}, nil
	case exp.Bool != nil:
		return &Bool{
			value: *exp.Bool,
		}, nil
	case exp.Identifier != nil:
		return p.Environment.GetVariable(*exp.Identifier)
	case exp.FunctionCall != nil:
		return p.EvalFunctionCall(exp.FunctionCall)
	}
	return nil, nil
}
