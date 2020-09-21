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

func (g *Generator) GenerateArrayIsEmpty(n string) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) IsEmpty() bool {\n", n))
	g.WriteString("return len(*t) == 0\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateMapIsEmpty(n string) {
	g.GenerateNosplit()
	g.WriteString(fmt.Sprintf("func (t *%s) IsEmpty() bool {\n", n))
	g.WriteString("return len(*t) == 0\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateStructAppendJsonString(n string, s *ast.StructType) {
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	if len(s.Fields.List) <= 0 {
		g.WriteString("return append(res, `{}`...)\n")
		g.WriteString("}\n\n")
	}

	g.WriteString("res = append(res, '{')\n")
	for _, f := range s.Fields.List {
		g.GenerateAppendJsonStringField(f) // `,` included
	}
	g.WriteString("res[len(res)-1] = '}'\n")
	g.WriteString("return res\n")
	g.WriteString("}\n\n")
}

func (g *Generator) GenerateMapAppendJsonString(n string, s *ast.MapType, c comment) {
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteString("if len(*t) <= 0 {\n")
	g.WriteString("return append(res, `[]`...)\n")
	g.WriteString("}\n")

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
	g.WriteString(fmt.Sprintf("func (t *%s) AppendJsonString(res []byte) []byte {\n", n))
	g.WriteString("if len(*t) <= 0 {\n")
	g.WriteString("return append(res, `[]`...)\n")
	g.WriteString("}\n")

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
		g.GenerateAppendInlinedInt("", access, j)
	case "int8":
		g.WriteString(fmt.Sprintf("res = lib.AppendSmallInt8(res, %s)\n", access))
	case "int16":
		g.GenerateAppendInlinedInt("16", access, j)
	case "int32":
		g.GenerateAppendInlinedInt("32", access, j)
	case "int64":
		g.GenerateAppendInlinedInt("64", access, j)
	case "uint":
		g.GenerateAppendInlinedUint("Uint", access, j)
	case "uint8":
		g.WriteString(fmt.Sprintf("res = lib.AppendSmallUint8(res, %s)\n", access))
	case "uint16":
		g.GenerateAppendInlinedUint("16", access, j)
	case "uint32":
		g.GenerateAppendInlinedUint("32", access, j)
	case "uint64":
		g.GenerateAppendInlinedUint("64", access, j)
	case "float32":
		g.WriteString(fmt.Sprintf("res = lib.AppendFloat32(res, %s, -1)\n", access))
	case "float64":
		g.WriteString(fmt.Sprintf("res = lib.AppendFloat64(res, %s, -1)\n", access))
	case "time.Time":
		if !isPointerAndNotOmitEmpty {
			access = "&" + access
		}

		if j.omitnano {
			g.WriteString(fmt.Sprintf("res = lib.AppendTimeWithoutNano(res, %s)\n", access))
		} else {
			g.WriteString(fmt.Sprintf("res = lib.AppendTime(res, %s)\n", access))
		}
	default:
		g.WriteString(fmt.Sprintf("res = %s.AppendJsonString(res)\n", strings.TrimPrefix(access, "*")))
	}

	if isPointerAndNotOmitEmpty {
		g.WriteString("}\n")
	}
}

func (g *Generator) GenerateAppendInlinedInt(size, access string, j jsonTag) {
	if j.small {
		if !j.unsigned {
			g.WriteString(fmt.Sprintf("if 0 <= %s {\n", access))
		}
		g.WriteString(fmt.Sprintf("res = lib.AppendSmallInt%s(res, %s)\n", size, access))
		if !j.unsigned {
			g.WriteString("} else {\n")
			g.WriteString(fmt.Sprintf("res = lib.AppendSmallMinusInt%s(res, %s)\n", size, access))
			g.WriteString("}\n")
		}
		return
	}

	if !j.unsigned {
		g.WriteString(fmt.Sprintf("if 0 <= %s {\n", access))
	}

	g.WriteString(fmt.Sprintf("if %s < lib.NSmalls {\n", access))
	g.WriteString(fmt.Sprintf("res = lib.AppendSmallInt%s(res, %s)\n", size, access))
	g.WriteString("} else {\n")
	g.WriteString(fmt.Sprintf("res = lib.AppendInt%s(res, %s)\n", size, access))
	g.WriteString("}\n")

	if !j.unsigned {
		g.WriteString("} else {\n")

		g.WriteString(fmt.Sprintf("if -lib.NSmalls < %s {\n", access))
		g.WriteString(fmt.Sprintf("res = lib.AppendSmallMinusInt%s(res, %s)\n", size, access))
		g.WriteString("} else {\n")
		g.WriteString(fmt.Sprintf("res = lib.AppendInt%s(res, %s)\n", size, access))
		g.WriteString("}\n")

		g.WriteString("}\n")
	}
}

func (g *Generator) GenerateAppendInlinedUint(size, access string, j jsonTag) {
	if j.small {
		g.WriteString(fmt.Sprintf("res = lib.AppendSmallUint%s(res, %s)\n", size, access))
		return
	}

	g.WriteString(fmt.Sprintf("if %s < lib.NSmalls {\n", access))
	g.WriteString(fmt.Sprintf("res = lib.AppendSmallUint%s(res, %s)\n", size, access))
	g.WriteString("} else {\n")
	g.WriteString(fmt.Sprintf("res = lib.AppendUint%s(res, %s)\n", size, access))
	g.WriteString("}\n")
}

var numReg = regexp.MustCompile("^u?int(?:8|16|32|64)?|float(?:32|64)$")

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
