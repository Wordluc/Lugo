package lugo

import (
	"Lugo/eval"
	"Lugo/parser"

	"github.com/alecthomas/participle/v2"
)

func GetProgram(code string) (*eval.Program, error) {
	parser, err := participle.Build[parser.Lua]()
	if err != nil {
		return nil, err
	}
	tr, err := parser.ParseString("program", code)
	if err != nil {
		return nil, err
	}
	return eval.NewEval(*tr), nil
}
