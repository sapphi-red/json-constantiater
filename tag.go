package main

import (
	"go/ast"
	"reflect"
	"strconv"
	"strings"
)

type fieldData struct {
	fieldName string
	tag       jsonTag
}

func getFieldData(f *ast.Field) (fd fieldData, skip bool) {
	fd.fieldName = f.Names[0].Name

	if f.Tag != nil {
		fd.tag = parseJsonTag(f.Tag.Value)
	}
	if fd.tag.name == "-" {
		return fd, true
	}

	if fd.tag.name == "" {
		fd.tag.name = ToSnakeCase(fd.fieldName)
	}

	return fd, false
}

type jsonTag struct {
	name      string
	noescape  bool
	omitempty bool
	omitnano  bool
}

func parseJsonTag(tag string) (j jsonTag) {
	tag, err := strconv.Unquote(tag)
	if err != nil {
		panic(err)
	}
	parsed := reflect.StructTag(tag)
	json := parsed.Get("json")
	return parseJsonTagInner(json)
}

func parseJsonTagInner(tag string) (j jsonTag) {
	strs := strings.Split(tag, ",")
	if len(strs) > 0 && strs[0] != "" {
		j.name = strs[0]
	}
	if len(strs) > 1 {
		for _, s := range strs {
			switch s {
			case "noescape":
				j.noescape = true
			case "omitempty":
				j.omitempty = true
			case "omitnano":
				j.omitnano = true
			}
		}
	}
	return
}
