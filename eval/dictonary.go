package eval

import "Lugo/parser"

func (p *Program) EvalTable(exp *parser.TableDeclaration) (res *Dictionary, err error) {
	res = &Dictionary{}
	res.Elements = make(map[Value]Value)
	for i, entry := range exp.Entries {

		v, e := p.EvalExp(*entry.Value)

		if e != nil {
			return res, e
		}

		if entry.Name != nil {
			var key Value
			if entry.Name.Identifier != nil {
				key = &String{value: *entry.Name.Identifier}
			} else {
				key, e = p.EvalValue(entry.Name)
			}
			if e != nil {
				return res, e
			}
			res.Elements[key] = v
		} else {
			key := Int{value: i}
			println(i)
			res.Elements[&key] = v
		}
	}
	return res, nil
}
