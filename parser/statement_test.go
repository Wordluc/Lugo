package parser

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
	tr, err := parser.ParseString("prova", "local a=3*(4+2)/4+1")
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex := "local a=(3 * (((4) + (2)) / (4))) + (1)\n"
	if res != ex {
		t.Fatalf("error %v expected: %v, got: %v", "simpleMathOperationWithLocal", ex, res)
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
	ex := "a=(3 * (4)) + (4)\n"
	if res != ex {
		t.Fatalf("error %v expected: %v, got: %v", "simpleMathWithoutLocal", ex, res)
	}
}
func TestMathWithMultAndDivide(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova", "a=3*4*4/4")
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex := "a=(3 * (4 * (4 / (4))))\n"
	if res != ex {
		t.Fatalf("error %v expected: %v, got: %v", "TestMathWithMultAndDivide", ex, res)
	}
}
func TestSimpleCallFuncWithParams(t *testing.T) {
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
		t.Fatalf("error %v expected: %v, got: %v", "CallFunction", ex, res)
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
	ex := "a=(2) + (2 * (4)) + (prova((2) + (3),(3),))\nlocal s=(3)\n"
	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "BlocStatement", ex, res)
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
		t.Fatalf("error %v expected: \n%v, got:\n %v", "AssignString", ex, res)
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
		t.Fatalf("error %v expected: \n%v, got:\n %v", "DeclarationFunction", ex+"|", res+"|")
	}
}
func TestFunctionDeclarationWithoutEnd(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	_, err = parser.ParseString("prova",
		`
	function prova()
	local a=3+4
	`)
	if err == nil {
		t.Fatal("should have 'end' at the end")
	}

}
func TestTableDeclaration(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova",
		`
		local a={
			"cioa",
			nome="luica"
		}
		local a=" fklsd" 
	`)
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex :=
		`local a={("cioa"),nome=("luica"),}
		local a=(" fklsd")`
	res = strings.ReplaceAll(res, "\u0009", "")
	ex = strings.ReplaceAll(ex, "\u0009", "")
	ex += "\n"

	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "DeclarationTable", ex+"|", res+"|")
	}
}
func TestTableRetrieve(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova",
		`
		local a=pippo.prova
		local a=pippo["cioa"]
		local a=pippo[getName()]
		local a=pippo[cioa]
		local a=pippo[3]
	`)
	if err != nil {
		t.Fatal(err.Error())
	}
	res := tr.toString()
	ex :=
		`local a=(pippo.prova)
		local a=(pippo[("cioa")])
		local a=(pippo[(getName())])
		local a=(pippo[(cioa)])
		local a=(pippo[(3)])`
	res = strings.ReplaceAll(res, "\u0009", "")
	ex = strings.ReplaceAll(ex, "\u0009", "")
	ex += "\n"

	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "RetrieveTable", ex+"|", res+"|")
	}
}
func TestLabdaFunction(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova",
		`
		prova=function(nome)
			local boo=nome
			local eta={12,43+3}
			return eta
		end
		return persona
	`)
	if err != nil {
		t.Fatal(err.Error())
	}
	res := tr.toString()
	ex :=
		`prova=function (nome,){local boo=(nome)
		local eta={(12),(43) + (3),}
		return (eta)
		}
		return (persona)`
	res = strings.ReplaceAll(res, "\u0009", "")
	ex = strings.ReplaceAll(ex, "\u0009", "")
	ex += "\n"

	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "LabdaFunction", ex+"|", res+"|")
	}
}
func TestTableWithLabdaFunction(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova",
		`
		local persona={
			nome="luca",
			eta=12,
			getFood=function()
				return "kebab" 
			end
		}
		return persona
	`)
	if err != nil {
		t.Fatal(err.Error())
	}
	res := tr.toString()
	ex :=
		`local persona={nome=("luca"),eta=(12),getFood=function (){
		return ("kebab")
		},}
		return (persona)`
	res = strings.ReplaceAll(res, "\u0009", "")
	ex = strings.ReplaceAll(ex, "\u0009", "")
	ex += "\n"

	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "LabdFunctionInTable", ex+"|", res+"|")
	}
}
func TestFunctionDeclarationWithMultipleParams(t *testing.T) {
	parser, err := participle.Build[Lua]()
	if err != nil {
		print(err.Error())
	}
	tr, err := parser.ParseString("prova",
		`
	function prova(a,b)
	local a=3>=4+4==4 + 4<4
	local b=34*4+4>3
	local f=3>=4==false
	prova()
	prova()
	end
	`)
	if err != nil {
		print(err.Error())
	}
	res := tr.toString()
	ex :=
		`function prova(a,b,){
		local a=(3 >= (4)) + (4 == (4)) + (4 < (4))
		local b=(34 * (4)) + (4 > (3))
		local f=(3 >= (4 == (false)))
		(prova())
		(prova())
	}`
	res = strings.ReplaceAll(res, "\u0009", "")
	ex = strings.ReplaceAll(ex, "\u0009", "")
	ex += "\n"

	if res != ex {
		t.Fatalf("error %v expected: \n%v, got:\n %v", "DeclarationFunction", ex+"|", res+"|")
	}
}
