package main

import (
	"Lugo/eval"
	"Lugo/parser"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
)

func main() {
	env := eval.NewEnvironment()
	env.AddCustomFunction("print", func(env *eval.Environment, args []eval.Value) eval.Value {
		for i := range args {
			fmt.Printf("%v ", args[i].Get())
		}
		println()
		return nil
	})
	env.AddCustomFunction("println", func(env *eval.Environment, args []eval.Value) eval.Value {
		for i := range args {
			fmt.Printf("%v ", args[i].Get())
		}
		println()
		println()
		return nil
	})

	var code string
	var multipleLine bool
	runCode := func(code *string) {
		pr, e := getProgram(*code, env)
		if e != nil {
			println("Error:", e.Error())
		}
		pr.Environment = env
		if e != nil {
			println("Error:", e.Error())
		}
		func() {
			e = pr.Run()
			if e != nil {
				println("Error:", e.Error())
			}
			e := recover()
			if e != nil {
				println("Error:", e.(string))
			}
		}()
		*code = ""

	}
	for {
		reader := bufio.NewReader(os.Stdin)
		if !multipleLine {
			fmt.Print("--")
		}
		t, _ := reader.ReadString('\n')
		if strings.Contains(t, "//") {
			multipleLine = !multipleLine
			runCode(&code)
			continue
		}

		code += t
		if !multipleLine {
			runCode(&code)
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
