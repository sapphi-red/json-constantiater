package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"

	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

/*
	compose exprs.

	```
	res = append(res, '{')
	res = append(res, `"`...)

	res = append(res, `"`...)
	res = append(res, '{')

	res = append(res, '{')
	res = append(res, '"')

	res = append(res, '"')
	res[len(res)-1] = '}'
	```

	into

	```
	res = append(res, `{"`...)
	res = append(res, `"}`...)
	res = append(res, `{"`...)
	res = append(res, '}')
	```
*/

func composeAppend(input string) []byte {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, "gen.go", input, parser.ParseComments)
	if err != nil {
		log.Fatalln(err, input)
	}

	newFile := astutil.Apply(file, func(cr *astutil.Cursor) bool {
		funcDecl, _ := cr.Node().(*ast.FuncDecl)
		if funcDecl == nil {
			return true
		}
		if funcDecl.Name.Name != "AppendJsonString" {
			return false
		}

		newBody := composeAppendBody(funcDecl.Body)
		funcDecl.Body = newBody
		cr.Replace(funcDecl)

		return false
	}, nil)

	var output bytes.Buffer
	err = format.Node(&output, fs, newFile)
	if err != nil {
		panic(err)
	}

	removedLineBreaks := bytes.Replace(
		bytes.Replace(
			bytes.Replace(output.Bytes(), []byte("\n\n\tres"), []byte("\n\tres"), -1),
			[]byte("\n\n\treturn res"),
			[]byte("\n\treturn res"),
			-1),
		[]byte("\n\n\tif"),
		[]byte("\n\tif"),
		-1,
	)

	return removedLineBreaks
}

func composeAppendBody(body *ast.BlockStmt) *ast.BlockStmt {
	newList := make([]ast.Stmt, 0, len(body.List))
	lits := make([]ast.BasicLit, 0)

	// `append` + `append`
	for _, stmt := range body.List {
		args := extractAppendFuncArg(stmt)
		if args != nil && len(*args) == 1 {
			arg := ifIsArgStringOrChar(*args)
			if arg != nil {
				lits = append(lits, *arg)
				continue
			}
		}

		if len(lits) > 0 {
			newList = appendComposedStmt(newList, lits)
			lits = make([]ast.BasicLit, 0)
		}

		newList = append(newList, stmt)
	}

	newerList := make([]ast.Stmt, 0, len(newList))
	newerList = append(newerList, newList[0])

	// `append` + `res[len(res)-1] =`
	for i := 1; i < len(newList); i++ {
		stmt := newList[i]

		assignLastLit := extractAssignLastLit(stmt)
		if assignLastLit == nil {
			newerList = append(newerList, stmt)
			continue
		}

		prevStmt := newList[i-1]
		args := extractAppendFuncArg(prevStmt)
		if args == nil || len(*args) != 1 {
			newerList = append(newerList, stmt)
			continue
		}

		arg := ifIsArgStringOrChar(*args)
		if arg == nil {
			newerList = append(newerList, stmt)
			continue
		}

		newerList[len(newerList)-1] = createReplacedStmt(arg, assignLastLit)
	}

	body.List = newerList
	return body
}

func appendComposedStmt(list []ast.Stmt, lits []ast.BasicLit) []ast.Stmt {
	newStr := ""
	for _, lit := range lits {
		litVal, _ := strconv.Unquote(lit.Value)
		newStr += litVal
	}

	var newLit *ast.BasicLit
	var ellipsis token.Pos
	if len(newStr) > 1 {
		newLit = &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(newStr),
		}
		ellipsis = 1
	} else {
		newLit = &ast.BasicLit{
			Kind:  token.CHAR,
			Value: strconv.QuoteRune(rune(newStr[0])),
		}
		ellipsis = 0
	}

	composed := ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent("res")},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: ast.NewIdent("append"),
			Args: []ast.Expr{
				ast.NewIdent("res"),
				newLit,
			},
			Ellipsis: ellipsis,
		}},
	}
	return append(list, &composed)
}

func createReplacedStmt(baseLit *ast.BasicLit, lastLit *ast.BasicLit) *ast.AssignStmt {
	base, _ := strconv.Unquote(baseLit.Value)
	last, _ := strconv.Unquote(lastLit.Value)

	newBase := base[:len(base)-1] + last

	var newLit *ast.BasicLit
	var ellipsis token.Pos
	if len(newBase) > 1 {
		newLit = &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(newBase),
		}
		ellipsis = 1
	} else {
		newLit = &ast.BasicLit{
			Kind:  token.CHAR,
			Value: strconv.QuoteRune(rune(newBase[0])),
		}
		ellipsis = 0
	}

	replaced := ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent("res")},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: ast.NewIdent("append"),
			Args: []ast.Expr{
				ast.NewIdent("res"),
				newLit,
			},
			Ellipsis: ellipsis,
		}},
	}
	return &replaced
}

func extractAppendFuncArg(stmt ast.Stmt) *[]ast.Expr {
	assignStmt, _ := stmt.(*ast.AssignStmt)
	if assignStmt == nil {
		return nil
	}

	lhs := assignStmt.Lhs
	if len(lhs) != 1 {
		return nil
	}
	lhId, _ := lhs[0].(*ast.Ident)
	if lhId == nil || lhId.Name != "res" {
		return nil
	}

	rhs := assignStmt.Rhs
	if len(rhs) != 1 {
		return nil
	}
	rhCall, _ := rhs[0].(*ast.CallExpr)
	if rhCall == nil {
		return nil
	}

	funcNameId, _ := rhCall.Fun.(*ast.Ident)
	if funcNameId == nil || funcNameId.Name != "append" {
		return nil
	}

	args := rhCall.Args[1:]
	return &args
}

func extractAssignLastLit(stmt ast.Stmt) *ast.BasicLit {
	assignStmt, _ := stmt.(*ast.AssignStmt)
	if assignStmt == nil {
		return nil
	}

	lhs := assignStmt.Lhs
	if len(lhs) != 1 {
		return nil
	}
	lhIndex, _ := lhs[0].(*ast.IndexExpr)
	if lhIndex == nil || !isLastIndexAccess(lhIndex) {
		return nil
	}

	rhs := assignStmt.Rhs
	if len(rhs) != 1 {
		return nil
	}

	rh := rhs[0]
	rhLit, _ := rh.(*ast.BasicLit)
	if rhLit == nil {
		return nil
	}
	if rhLit.Kind == token.CHAR {
		return rhLit
	}
	return nil
}

func isLastIndexAccess(indexExpr *ast.IndexExpr) bool {
	sliceId, _ := indexExpr.X.(*ast.Ident)
	if sliceId == nil || sliceId.Name != "res" {
		return false
	}

	lhIndexBinary, _ := indexExpr.Index.(*ast.BinaryExpr)
	if lhIndexBinary == nil || lhIndexBinary.Op != token.SUB {
		return false
	}
	lhIndexCall, _ := lhIndexBinary.X.(*ast.CallExpr)
	lhIndexOne, _ := lhIndexBinary.Y.(*ast.BasicLit)
	if lhIndexCall == nil || lhIndexOne == nil || lhIndexOne.Value != "1" {
		return false
	}
	lhIndexCallFnId, _ := lhIndexCall.Fun.(*ast.Ident)
	if lhIndexCallFnId == nil || lhIndexCallFnId.Name != "len" {
		return false
	}

	return true
}

func ifIsArgStringOrChar(args []ast.Expr) *ast.BasicLit {
	arg := args[0]
	litArg, _ := arg.(*ast.BasicLit)
	if litArg == nil {
		return nil
	}
	if litArg.Kind == token.STRING || litArg.Kind == token.CHAR {
		return litArg
	}
	return nil
}
