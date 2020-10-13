package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"strings"

	"golang.org/x/tools/go/ast/inspector"
)

var input = flag.String("input", "main.go", "input file name")
var output = flag.String("output", "", "output file name")

func main() {
	flag.Parse()

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, *input, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	g := Generator{}
	g.GenerateHead()
	g.DeclarePkgNameAndImports(f.Name.Name)

	inspect := inspector.New([]*ast.File{f})
	filter := []ast.Node{new(ast.GenDecl)}
	inspect.Preorder(filter, func(n ast.Node) {
		genDecl := n.(*ast.GenDecl)
		if genDecl.Tok != token.TYPE || len(genDecl.Specs) > 1 {
			return
		}

		typDef := genDecl.Specs[0].(*ast.TypeSpec)
		name := typDef.Name.Name
		comment := parseComments(genDecl.Doc)

		switch typ := typDef.Type.(type) {
		case *ast.StructType:
			g.GenerateNewJsonMarshal(name)
			g.GenerateWriteJsonString(name)
			g.GenerateStructAppendJsonString(name, typ)
		case *ast.ArrayType:
			g.GenerateNewJsonMarshal(name)
			g.GenerateWriteJsonString(name)
			g.GenerateArrayAppendJsonString(name, typ, comment)
			g.GenerateArrayIsEmpty(name)
		case *ast.MapType:
			if types.ExprString(typ.Key) == "string" {
				g.GenerateNewJsonMarshal(name)
				g.GenerateWriteJsonString(name)
				g.GenerateMapAppendJsonString(name, typ, comment)
				g.GenerateMapIsEmpty(name)
			}
		}
	})

	src := g.FormatGetString()

	if *output == "" {
		*output = strings.TrimSuffix(*input, ".go") + "_constantiated.go"
	}

	err = ioutil.WriteFile(*output, src, 0644)
	if err != nil {
		panic(err)
	}
}
