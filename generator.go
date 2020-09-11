package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/types"
	"regexp"
	"strings"

	"github.com/sapphi-red/json-constantiater/lib"
)

type Generator struct {
	bytes.Buffer
}

func (g *Generator) GenerateHead() {
	g.WriteString("// Code generated by json-constantiater DO NOT EDIT\n")
	g.WriteString("\n")
}

func (g *Generator) FormatGetString() []byte {
	composed := composeAppend(g.String())
	return composed
}

func (g *Generator) DeclarePkgNameAndImports(name string) {
	g.WriteString(fmt.Sprintf("package %s\n\n", name))
	g.WriteString("import \"github.com/sapphi-red/json-constantiater/lib\"\n\n")
}

func (g *Generator) GenerateNosplit() {
	g.WriteString("//go:nosplit\n")
}

func (g *Generator) GenerateNewJsonMarshal(n string) {
	g.WriteString(fmt.Sprintf("func (t *%s) NewJsonMarshal() []byte {\n", n))
	g.WriteString("tmpPtr := lib.GetFromPool()\n")
	g.WriteString("tmp := *tmpPtr\n")
	g.WriteString("tmp = t.AppendJsonString(tmp)\n")
	g.WriteString("res := make([]byte, len(tmp))\n")
	g.WriteString("copy(res, tmp)\n")
	g.WriteString("*tmpPtr = tmp\n")
	g.WriteString("lib.PutToPool(tmpPtr)\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateStructJsonLen(n string, s *ast.StructType) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) JsonLen() int {\n", n))
	g.WriteString("l := 2\n") // head `{` and tail `}`
	for _, f := range s.Fields.List {
		g.GenerateJsonLenField(f) // `,` included
	}
	g.WriteString("return l - 1\n") // 1 for tail `,`
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateArrayJsonLen(n string, s *ast.ArrayType, c comment) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) JsonLen() int {\n", n))
	g.WriteString("l := 2\n") // head `[` and tail `]`

	g.WriteString("for _, e := range *t {\n")
	g.GenerateJsonLenSingle("e", s.Elt, c.value)
	g.WriteString("l += 1\n") // 1 for `,`
	g.WriteString("}\n")

	g.WriteString("return l - 1\n") // 1 for tail `,`
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateMapJsonLen(n string, s *ast.MapType, c comment) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) JsonLen() int {\n", n))
	g.WriteString("l := 2\n") // head `{` and tail `}`

	g.WriteString("for k, v := range *t {\n")
	g.GenerateJsonLenSingle("k", s.Key, c.key)
	g.GenerateJsonLenSingle("v", s.Value, c.value)
	g.WriteString("l += 1 + 1\n") // 1 for `:`, 1 for `,`
	g.WriteString("}\n")

	g.WriteString("return l - 1\n") // 1 for tail `,`
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateArrayIsEmpty(n string) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) IsEmpty() bool {\n", n))
	g.WriteString("return len(*t) == 0")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateMapIsEmpty(n string) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) IsEmpty() bool {\n", n))
	g.WriteString("return len(*t) == 0")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateStructAppendJsonString(n string, s *ast.StructType) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteString("res = append(res, '{')\n")
	for _, f := range s.Fields.List {
		g.GenerateAppendJsonStringField(f) // `,` included
	}
	g.WriteString("res[len(res)-1] = '}'\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateMapAppendJsonString(n string, s *ast.MapType, c comment) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteString("res = append(res, '{')\n")

	g.WriteString("for k, v := range *t {\n")
	g.GenerateAppendJsonStringValue("k", s.Key, c.key)
	g.WriteString("res = append(res, ':')\n")
	g.GenerateAppendJsonStringValue("v", s.Value, c.value)
	g.WriteString("res = append(res, ',')\n")
	g.WriteString("}\n")

	g.WriteString("res[len(res)-1] = '}'\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateArrayAppendJsonString(n string, s *ast.ArrayType, c comment) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteString("res = append(res, '[')\n")

	g.WriteString("for _, e := range *t {\n")
	g.GenerateAppendJsonStringValue("e", s.Elt, c.value)
	g.WriteString("res = append(res, ',')\n")
	g.WriteString("}\n")

	g.WriteString("res[len(res)-1] = ']'\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
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

	g.WriteString("res = append(res, ',')\n")

	if j.omitempty {
		g.WriteString("}\n")
	}
}

func (g *Generator) GenerateAppendJsonStringValue(access string, typeExpr ast.Expr, j jsonTag) {
	typName := types.ExprString(typeExpr)
	isPointerAndNotOmitEmpty := !j.omitempty && strings.HasPrefix(typName, "*")

	if isPointerAndNotOmitEmpty {
		g.WriteString(fmt.Sprintf("if %s == nil {\n", access))
		g.WriteString("res = append(res, `null`...)\n")
		g.WriteString("} else {\n")

		typName = strings.TrimPrefix(typName, "*")
		access = "*" + access
	}

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
	case "float32":
		g.WriteString(fmt.Sprintf("res = lib.AppendFloat32(res, %s, -1)\n", access))
	case "float64":
		g.WriteString(fmt.Sprintf("res = lib.AppendFloat64(res, %s, -1)\n", access))
	case "time.Time":
		if isPointerAndNotOmitEmpty {
			g.WriteString(fmt.Sprintf("res = lib.AppendTime(res, %s)\n", access))
		} else {
			g.WriteString(fmt.Sprintf("res = lib.AppendTime(res, &%s)\n", access))
		}
	default:
		g.WriteString(fmt.Sprintf("res = %s.AppendJsonString(res)\n", strings.TrimPrefix(access, "*")))
	}

	if isPointerAndNotOmitEmpty {
		g.WriteString("}\n")
	}
}

// includes tail `,`
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
	g.WriteString(fmt.Sprintf("l += 2 + %d + 1 + 1\n", lib.GetEscapedLen(fd.tag.name))) // 2 for `"`, 1 for `:`, 1 for tail `,`
	if j.omitempty {
		g.WriteString("}\n")
	}
}

var numReg = regexp.MustCompile("^u?int(?:8|16|32|64)?|float(?:32|64)$")

func (g *Generator) GenerateJsonLenSingle(access string, typeExpr ast.Expr, j jsonTag) {
	typName := types.ExprString(typeExpr)
	isPointerAndNotOmitEmpty := !j.omitempty && strings.HasPrefix(typName, "*")

	if isPointerAndNotOmitEmpty {
		g.WriteString(fmt.Sprintf("if %s == nil {\n", access))
		g.WriteString("l += 4\n") // 4 for null
		g.WriteString("} else {\n")

		typName = strings.TrimPrefix(typName, "*")
		access = "*" + access
	}

	if j.len > 0 {
		if typName == "string" {
			g.WriteString(fmt.Sprintf("l += %d\n", 2+j.len)) // 2 for `"`
		} else {
			g.WriteString(fmt.Sprintf("l += %d\n", j.len))
		}
		return
	}

	switch typName {
	case "string":
		g.WriteString("l += 2 + ") // 2 for `"`
		if j.noescape {
			g.WriteString(fmt.Sprintf("len(%s)\n", access))
		} else {
			g.WriteString(fmt.Sprintf("lib.GetEscapedLen(%s)\n", access))
		}
	case "bool":
		g.WriteString(fmt.Sprintf("l += %d\n", getLenOfSimpleType(typName)))
	case "time.Time":
		g.WriteString("l += 32\n")
	default:
		if numReg.MatchString(typName) {
			g.WriteString(fmt.Sprintf("l += %d\n", getLenOfSimpleType(typName)))
		} else {
			g.WriteString(fmt.Sprintf("l += %s.JsonLen()\n", strings.TrimPrefix(access, "*")))
		}
	}

	if isPointerAndNotOmitEmpty {
		g.WriteString("}\n")
	}
}

func (g *Generator) GenerateOmitEmptyIfNot(access string, typeExpr ast.Expr) {
	typName := types.ExprString(typeExpr)
	switch typName {
	case "string":
		g.WriteString(fmt.Sprintf("if %s != \"\" {\n", access))
	case "bool":
		g.WriteString(fmt.Sprintf("if %s {\n", access))
	case "time.Time":
		g.WriteString(fmt.Sprintf("if !%s.IsZero() {\n", access))
	default:
		if numReg.MatchString(typName) {
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
