package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/types"
	"regexp"
	"strings"
)

type Generator struct {
	bytes.Buffer
}

func (g *Generator) WriteLine(s string) {
	g.WriteString(s)
	g.WriteBR()
}

func (g *Generator) WriteLinef(format string, a ...interface{}) {
	g.WriteLine(fmt.Sprintf(format, a...))
}

func (g *Generator) WriteBR() {
	g.WriteByte('\n')
}

func (g *Generator) GenerateHead() {
	g.WriteLine("// Code generated by json-constantiater DO NOT EDIT")
	g.WriteBR()
}

func (g *Generator) FormatGetString() []byte {
	composed := composeAppend(g.String())
	return composed
}

func (g *Generator) DeclarePkgNameAndImports(name string) {
	g.WriteLinef("package %s", name)
	g.WriteBR()
	g.WriteLine(`import "github.com/sapphi-red/json-constantiater/lib"`)
	g.WriteBR()
}

func (g *Generator) GenerateNosplit() {
	g.WriteLine("//go:nosplit")
}

func (g *Generator) GenerateNewJsonMarshal(n string) {
	g.WriteLinef("func (t *%s) NewJsonMarshal() []byte {", n)
	g.WriteLine("tmpPtr := lib.GetFromPool()")
	g.WriteLine("tmp := *tmpPtr")
	g.WriteLine("tmp = t.AppendJsonString(tmp)")
	g.WriteLine("res := make([]byte, len(tmp))")
	g.WriteLine("copy(res, tmp)")
	g.WriteLine("*tmpPtr = tmp")
	g.WriteLine("lib.PutToPool(tmpPtr)")
	g.WriteLine("return res")
	g.WriteLine("}")
	g.WriteBR()
}

func (g *Generator) GenerateArrayIsEmpty(n string) {
	g.GenerateNosplit()
	g.WriteLinef("func (t *%s) IsEmpty() bool {", n)
	g.WriteLine("return len(*t) == 0")
	g.WriteLine("}")
	g.WriteBR()
}

func (g *Generator) GenerateMapIsEmpty(n string) {
	g.GenerateNosplit()
	g.WriteLinef("func (t *%s) IsEmpty() bool {", n)
	g.WriteLine("return len(*t) == 0")
	g.WriteLine("}")
	g.WriteBR()
}

func (g *Generator) GenerateStructAppendJsonString(n string, s *ast.StructType) {
	g.WriteLinef("func (t *%s) AppendJsonString(res []byte) []byte {", n)
	if len(s.Fields.List) <= 0 {
		g.WriteLine("return append(res, `{}`...)")
		g.WriteLine("}")
		g.WriteBR()
	}

	g.WriteLine("res = append(res, '{')")
	for _, f := range s.Fields.List {
		g.GenerateAppendJsonStringField(f) // `,` included
	}
	g.WriteLine("res[len(res)-1] = '}'")
	g.WriteLine("return res")
	g.WriteLine("}")
	g.WriteBR()
}

func (g *Generator) GenerateMapAppendJsonString(n string, s *ast.MapType, c comment) {
	g.WriteLinef("func (t *%s) AppendJsonString(res []byte) []byte {", n)
	g.WriteLine("if len(*t) <= 0 {")
	g.WriteLine("return append(res, `[]`...)")
	g.WriteLine("}")

	g.WriteLine("res = append(res, '{')")

	g.WriteLine("for k, v := range *t {")
	g.GenerateAppendJsonStringValue("k", s.Key, c.key)
	g.WriteLine("res = append(res, ':')")
	g.GenerateAppendJsonStringValue("v", s.Value, c.value)
	g.WriteLine("res = append(res, ',')")
	g.WriteLine("}")

	g.WriteLine("res[len(res)-1] = '}'")
	g.WriteLine("return res")
	g.WriteLine("}")
	g.WriteBR()
}

func (g *Generator) GenerateArrayAppendJsonString(n string, s *ast.ArrayType, c comment) {
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteLine("if len(*t) <= 0 {")
	g.WriteLine("return append(res, `[]`...)")
	g.WriteLine("}")

	g.WriteLine("res = append(res, '[')")

	g.WriteLine("for _, e := range *t {")
	g.GenerateAppendJsonStringValue("e", s.Elt, c.value)
	g.WriteLine("res = append(res, ',')")
	g.WriteLine("}")

	g.WriteLine("res[len(res)-1] = ']'")
	g.WriteLine("return res")
	g.WriteLine("}")
	g.WriteBR()
}

// includes tail `,`
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

	g.WriteString(fmt.Sprintf("res = append(res, `\"%s\":`...)\n", escapeString(j.name)))
	g.GenerateAppendJsonStringValue(access, f.Type, j)

	g.WriteLine("res = append(res, ',')")

	if j.omitempty {
		g.WriteLine("}")
	}
}

func (g *Generator) GenerateAppendJsonStringValue(access string, typeExpr ast.Expr, j jsonTag) {
	typName := types.ExprString(typeExpr)
	isPointerAndNotOmitEmpty := !j.omitempty && strings.HasPrefix(typName, "*")

	if isPointerAndNotOmitEmpty {
		g.WriteString(fmt.Sprintf("if %s == nil {\n", access))
		g.WriteLine("res = append(res, `null`...)")
		g.WriteLine("} else {")

		typName = strings.TrimPrefix(typName, "*")
		access = "*" + access
	}

	switch typName {
	case "string":
		g.WriteLine(`res = append(res, '"')`)
		if j.noescape {
			g.WriteLinef("res = append(res, %s...)", access)
		} else {
			g.WriteLinef("res = lib.AppendByteWithEscape(res, %s)", access)
		}
		g.WriteLine("res = append(res, '\"')")
	case "bool":
		g.WriteLinef("res = lib.AppendBool(res, %s)", access)
	case "int":
		g.GenerateAppendInlinedInt("", access, j)
	case "int8":
		g.WriteLinef("res = lib.AppendSmallInt8(res, %s)", access)
	case "int16":
		g.GenerateAppendInlinedInt("16", access, j)
	case "int32":
		g.GenerateAppendInlinedInt("32", access, j)
	case "int64":
		g.GenerateAppendInlinedInt("64", access, j)
	case "uint":
		g.GenerateAppendInlinedUint("Uint", access, j)
	case "uint8":
		g.WriteLinef("res = lib.AppendSmallUint8(res, %s)", access)
	case "uint16":
		g.GenerateAppendInlinedUint("16", access, j)
	case "uint32":
		g.GenerateAppendInlinedUint("32", access, j)
	case "uint64":
		g.GenerateAppendInlinedUint("64", access, j)
	case "float32":
		g.WriteLinef("res = lib.AppendFloat32(res, %s, -1)", access)
	case "float64":
		g.WriteLinef("res = lib.AppendFloat64(res, %s, -1)", access)
	case "time.Time":
		if !isPointerAndNotOmitEmpty {
			access = "&" + access
		}

		if j.omitnano {
			g.WriteLinef("res = lib.AppendTimeWithoutNano(res, %s)", access)
		} else {
			g.WriteLinef("res = lib.AppendTime(res, %s)", access)
		}
	default:
		g.WriteLinef("res = %s.AppendJsonString(res)", strings.TrimPrefix(access, "*"))
	}

	if isPointerAndNotOmitEmpty {
		g.WriteLine("}")
	}
}

func (g *Generator) GenerateAppendInlinedInt(size, access string, j jsonTag) {
	if j.small {
		if !j.unsigned {
			g.WriteLinef("if 0 <= %s {", access)
		}
		g.WriteLinef("res = lib.AppendSmallInt%s(res, %s)", size, access)
		if !j.unsigned {
			g.WriteLine("} else {")
			g.WriteLinef("res = lib.AppendSmallMinusInt%s(res, %s)", size, access)
			g.WriteLine("}")
		}
		return
	}

	if !j.unsigned {
		g.WriteLinef("if 0 <= %s {", access)
	}

	g.WriteLinef("if %s < lib.NSmalls {", access)
	g.WriteLinef("res = lib.AppendSmallInt%s(res, %s)", size, access)
	g.WriteLine("} else {")
	g.WriteLinef("res = lib.AppendInt%s(res, %s)", size, access)
	g.WriteLine("}")

	if !j.unsigned {
		g.WriteLine("} else {")

		g.WriteLinef("if -lib.NSmalls < %s {", access)
		g.WriteLinef("res = lib.AppendSmallMinusInt%s(res, %s)", size, access)
		g.WriteLine("} else {")
		g.WriteLinef("res = lib.AppendInt%s(res, %s)", size, access)
		g.WriteLine("}")

		g.WriteLine("}")
	}
}

func (g *Generator) GenerateAppendInlinedUint(size, access string, j jsonTag) {
	if j.small {
		g.WriteLinef("res = lib.AppendSmallUint%s(res, %s)", size, access)
		return
	}

	g.WriteLinef("if %s < lib.NSmalls {", access)
	g.WriteLinef("res = lib.AppendSmallUint%s(res, %s)", size, access)
	g.WriteLine("} else {")
	g.WriteLinef("res = lib.AppendUint%s(res, %s)", size, access)
	g.WriteLine("}")
}

var numReg = regexp.MustCompile("^u?int(?:8|16|32|64)?|float(?:32|64)$")

func (g *Generator) GenerateOmitEmptyIfNot(access string, typeExpr ast.Expr) {
	typName := types.ExprString(typeExpr)
	switch typName {
	case "string":
		g.WriteLinef("if %s != \"\" {\n", access)
	case "bool":
		g.WriteLinef("if %s {\n", access)
	case "time.Time":
		g.WriteLinef("if !%s.IsZero() {\n", access)
	default:
		if numReg.MatchString(typName) {
			g.WriteLinef("if %s != 0 {\n", access)
		} else if strings.HasPrefix(typName, "[]") || strings.HasPrefix(typName, "map") {
			g.WriteLinef("if len(%s) > 0 {\n", access)
		} else if strings.HasPrefix(typName, "*") {
			g.WriteLinef("if %s != nil {\n", access)
		} else {
			g.WriteLinef("if !%s.IsEmpty() {\n", access)
		}
	}
}
