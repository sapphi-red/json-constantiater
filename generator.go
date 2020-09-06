package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/types"
	"regexp"
	"strings"
)

type Generator struct {
	bytes.Buffer
}

func (g *Generator) GenerateHead() {
	g.WriteString("// Code generated by json-constantiater DO NOT EDIT\n")
	g.WriteString("\n")
}

func (g *Generator) FormatGetString() []byte {
	src, err := format.Source(g.Bytes())
	if err != nil {
		panic(err)
	}
	return src
}

func (g *Generator) DeclarePkgNameAndImports(name string) {
	g.WriteString(fmt.Sprintf("package %s\n\n", name))
	g.WriteString("import \"github.com/sapphi-red/json-constantiater/lib\"\n\n")
}

func (g *Generator) GenerateNosplit() {
	g.WriteString("//go:nosplit\n")
}

func (g *Generator) GenerateNewJsonMarshal(n string) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) NewJsonMarshal() []byte {\n", n))
	g.WriteString("res := make([]byte, 0, t.JsonLen())\n")
	g.WriteString("res = t.AppendJsonString(res)\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateStructJsonLen(n string, s *ast.StructType) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) JsonLen() int64 {\n", n))
	g.WriteString("var l int64 = 0\n")
	for _, f := range s.Fields.List {
		g.GenerateJsonLenField(f)
	}
	g.WriteString("return l\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateArrayJsonLen(n string, s *ast.ArrayType) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) JsonLen() int64 {\n", n))
	g.WriteString("var l int64 = 0\n")

	g.WriteString("for _, e := range *t {\n")
	g.WriteString("l += e.JsonLen()\n")
	g.WriteString("}\n")

	g.WriteString("return l\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateMapJsonLen(n string, s *ast.MapType) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) JsonLen() int64 {\n", n))
	g.WriteString("var l int64 = 0\n")

	g.WriteString("for _, e := range *t {\n")
	g.WriteString("l += e.JsonLen()\n")
	g.WriteString("}\n")

	g.WriteString("return l\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateStructAppendJsonString(n string, s *ast.StructType) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteString("res = append(res, '{')\n")
	for i, f := range s.Fields.List {
		g.GenerateAppendJsonStringField(f)
		if i != len(s.Fields.List) - 1 {
			g.WriteString("res = append(res, ',')\n")
		}
	}
	g.WriteString("res = append(res, '}')\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateMapAppendJsonString(n string, s *ast.MapType) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteString("res = append(res, '[')\n")

	g.WriteString("for k, v := range *t {\n")
	g.GenerateAppendJsonStringValue("k", s.Key, jsonTag{})
	g.WriteString("res = append(res, ',')\n")
	g.GenerateAppendJsonStringValue("v", s.Value, jsonTag{})
	g.WriteString("res = append(res, ',')\n")
	g.WriteString("}\n")

	g.WriteString("res = res[:len(res)-1]\n")
	g.WriteString("res = append(res, ']')\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateArrayAppendJsonString(n string, s *ast.ArrayType) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteString("res = append(res, '[')\n")

	g.WriteString("for _, e := range *t {\n")
	g.GenerateAppendJsonStringValue("e", s.Elt, jsonTag{})
	g.WriteString("res = append(res, ',')\n")
	g.WriteString("}\n")

	g.WriteString("res = res[:len(res)-1]\n")
	g.WriteString("res = append(res, ']')\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateAppendJsonStringField(f *ast.Field) {
	if len(f.Names) > 1 {
		panic("doesnt support several names")
	}
	fd, skip := getFieldData(f)
	if skip {
		return
	}

	access := "t." + fd.fieldName
	j := fd.tag

	if j.omitempty {
		g.GenerateOmitEmptyIfNot(access, f.Type)
	}

	g.WriteString(fmt.Sprintf("res = append(res, `\"%s\":`...)\n", j.name))
	g.GenerateAppendJsonStringValue(access, f.Type, j)

	if j.omitempty {
		g.WriteString("}\n")
	}
}

func (g *Generator) GenerateAppendJsonStringValue(access string, typeExpr ast.Expr, j jsonTag) {
	typName := types.ExprString(typeExpr)
	switch typName {
	case "string":
		g.WriteString("res = append(res, '\"')\n")
		if j.noescape {
			g.WriteString(fmt.Sprintf("res = append(res, %s...)\n", access))
		} else {
			g.WriteString(fmt.Sprintf("res = lib.AppendByteWithEscape(res, %s)\n", access))
		}
		g.WriteString("res = append(res, '\"')\n")
	case "bool":
		g.WriteString(fmt.Sprintf("res = lib.AppendBool(res, %s)\n", access))
	case "int":
		g.WriteString(fmt.Sprintf("res = lib.AppendInt(res, %s)\n", access))
	case "int8":
		g.WriteString(fmt.Sprintf("res = lib.AppendInt8(res, %s)\n", access))
	case "int16":
		g.WriteString(fmt.Sprintf("res = lib.AppendInt16(res, %s)\n", access))
	case "int32":
		g.WriteString(fmt.Sprintf("res = lib.AppendInt32(res, %s)\n", access))
	case "int64":
		g.WriteString(fmt.Sprintf("res = lib.AppendInt64(res, %s)\n", access))
	case "uint":
		g.WriteString(fmt.Sprintf("res = lib.AppendUint(res, %s)\n", access))
	case "uint8":
		g.WriteString(fmt.Sprintf("res = lib.AppendUint8(res, %s)\n", access))
	case "uint16":
		g.WriteString(fmt.Sprintf("res = lib.AppendUint16(res, %s)\n", access))
	case "uint32":
		g.WriteString(fmt.Sprintf("res = lib.AppendUint32(res, %s)\n", access))
	case "uint64":
		g.WriteString(fmt.Sprintf("res = lib.AppendUint64(res, %s)\n", access))
	default:
		g.WriteString(fmt.Sprintf("res = %s.AppendJsonString(res)\n", access))
	}
}

func (g *Generator) GenerateJsonLenField(f *ast.Field) {
	if len(f.Names) > 1 {
		panic("doesnt support several names")
	}

	fd, skip := getFieldData(f)
	if skip {
		return
	}

	access := "t." + fd.fieldName
	j := fd.tag

	if j.omitempty {
		g.GenerateOmitEmptyIfNot(access, f.Type)
	}
	g.GenerateJsonLenSingle(access, f.Type, j)
	if j.omitempty {
		g.WriteString("}\n")
	}
}

var intReg = regexp.MustCompile("^u?int(?:8|16|32|64)?$")

func (g *Generator) GenerateJsonLenSingle(access string, typeExpr ast.Expr, j jsonTag) {
	typName := types.ExprString(typeExpr)
	switch typName {
	case "string":
		g.WriteString("l += ")
		if j.noescape {
			g.WriteString(fmt.Sprintf("int64(len(%s))\n", access))
		} else {
			g.WriteString(fmt.Sprintf("lib.GetEscapedLen(%s)\n", access))
		}
	case "bool":
		g.WriteString(fmt.Sprintf("l += %d\n", getLenOfSimpleType(typName)))
	default:
		if intReg.MatchString(typName) {
			g.WriteString(fmt.Sprintf("l += %d\n", getLenOfSimpleType(typName)))
		} else {
			g.WriteString(fmt.Sprintf("l += %s.JsonLen()\n", access))
		}
	}
}

func (g *Generator) GenerateOmitEmptyIfNot(access string, typeExpr ast.Expr) {
	typName := types.ExprString(typeExpr)
	switch typName {
	case "string":
		g.WriteString(fmt.Sprintf("if %s != \"\" {\n", access))
	case "bool":
		g.WriteString(fmt.Sprintf("if %s {\n", access))
	default:
		if intReg.MatchString(typName) {
			g.WriteString(fmt.Sprintf("if %s != 0 {\n", access))
		} else if strings.HasPrefix(typName, "[]") || strings.HasPrefix(typName, "map") {
			g.WriteString(fmt.Sprintf("if len(%s) > 0 {\n", access))
		} else if strings.HasPrefix(typName, "*") {
			g.WriteString(fmt.Sprintf("if %s != nil {\n", access))
		} else {
			g.WriteString(fmt.Sprintf("if !%s.IsEmpty() {\n", access))
		}
	}
}
