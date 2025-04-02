package grammar

import (
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
	ex := "local a=(3 * 4) + (2 / 4) + (1)"
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
	ex := "a=(3 * 4) + (4)"
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
	ex := "a=(prova((2) + (3),(3),))"
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
	ex := "a=(2) + (2 * 4) + (prova((2) + (3),(3),))\nlocal s=(3)"
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
	ex := `a=("ciao")`
	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "simpleMath", ex, res)
	}
}
