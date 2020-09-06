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
	f, err := parser.ParseFile(fs, *input, nil, 0)
	if err != nil {
		panic(err)
	}

	g := Generator{}
	g.GenerateHead()
	g.DeclarePkgNameAndImports(f.Name.Name)

	inspect := inspector.New([]*ast.File{f})
	filter := []ast.Node{new(ast.TypeSpec)}
	inspect.Preorder(filter, func(n ast.Node) {
		typ := n.(*ast.TypeSpec)
		name := typ.Name.Name

		switch typ := typ.Type.(type) {
		case *ast.StructType:
			g.GenerateNewJsonMarshal(name)
			g.GenerateStructAppendJsonString(name, typ)
			g.GenerateStructJsonLen(name, typ)
		case *ast.ArrayType:
			g.GenerateNewJsonMarshal(name)
			g.GenerateArrayAppendJsonString(name, typ)
			g.GenerateArrayJsonLen(name, typ)
		case *ast.MapType:
			if types.ExprString(typ.Key) == "string" {
				g.GenerateNewJsonMarshal(name)
				g.GenerateMapAppendJsonString(name, typ)
				g.GenerateMapJsonLen(name, typ)
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
