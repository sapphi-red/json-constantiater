package main

import (
	"regexp"
	"strings"

	"github.com/sapphi-red/json-constantiater/lib"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func escapeString(str string) string {
	dest := make([]byte, 0, len(str))
	dest = lib.AppendByteWithEscape(dest, &str)
	return string(dest)
}
