package main

import (
	"Lugo/eval"
	"Lugo/parser"
	"bufio"
	"fmt"
	"os"

	"github.com/alecthomas/participle/v2"
)

func main() {
	env := eval.NewEnvironment()
	env.AddCustomFunction("print", func(env *eval.Environment) eval.Value {
		v, e := env.GetVariable("0")
		if e != nil {
			println(e.Error())
			return nil
		}
		println(v.(*eval.String).Get())
		return nil
	})
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("--")
		code, _ := reader.ReadString('\n')
		pr, e := getProgram(code, env)
		if e != nil {
			println("Error:", e.Error())
		}
		pr.Environment = env
		if e != nil {
			println("Error:", e.Error())
		}
		e = pr.Run()
		if e != nil {
			println("Error:", e.Error())
		}
	}

}
func getProgram(code string, env *eval.Environment) (*eval.Program, error) {
	parser, err := participle.Build[parser.Lua]()
	if err != nil {
		return nil, err
	}
	tr, err := parser.ParseString("program", code)
	if err != nil {
		return nil, err
	}
	return eval.NewCustomEval(*tr, env), nil
}
