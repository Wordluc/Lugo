package grammar

import (
	"strings"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestSimpleMathOperationWithLocal(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova", "local a=3*4+2/4+1")
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex := "local a=(3 * 4) + (2 / 4) + (1)\n"
	if res != ex {
		t.Fatalf("error %v expected: %v, got: %v", "simpleMath", ex, res)
	}
}
func TestSimpleMathOperationWithoutLocal(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova", "a=3*4+4")
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex := "a=(3 * 4) + (4)\n"
	if res != ex {
		t.Fatalf("error %v expected: %v, got: %v", "simpleMath", ex, res)
	}
}
func TestSimpleCallFuncWithParms(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova", "a = prova(2+3,3)")
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex := "a=(prova((2) + (3),(3),))\n"
	if res != ex {
		t.Fatalf("error %v expected: %v, got: %v", "simpleMath", ex, res)
	}
}
func TestBlockStatement(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova", "a =2 +2*4+ prova(2+3,3)\nlocal s=3")
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex := "a=(2) + (2 * 4) + (prova((2) + (3),(3),))\nlocal s=(3)\n"
	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "simpleMath", ex, res)
	}
}
func TestString(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova", `a = "ciao"`)
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex := `a=("ciao")` + "\n"
	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "simpleMath", ex, res)
	}
}
func TestFunctionDeclaration(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova",
		`
	function prova()
	local a=3+4
	local b=3+4
	prova()
	prova()
	end
	`)
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex :=
		`function prova(){
		local a=(3) + (4)
		local b=(3) + (4)
		(prova())
		(prova())
	}`
	res = strings.ReplaceAll(res, "\u0009", "")
	ex = strings.ReplaceAll(ex, "\u0009", "")
	ex += "\n"

	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "simpleMath", ex+"|", res+"|")
	}
}
