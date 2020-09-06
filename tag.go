package main

import (
	"reflect"
	"strconv"
	"strings"
)

type jsonTag struct {
	name      string
	noescape  bool
	omitempty bool
}

func parseJsonTag(tag string) (j jsonTag) {
	tag, err := strconv.Unquote(tag)
	if err != nil {
		panic(err)
	}
	parsed := reflect.StructTag(tag)
	json := parsed.Get("json")
	strs := strings.Split(json, ",")
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
			}
		}
	}
	return j
}
