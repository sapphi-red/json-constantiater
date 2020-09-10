package main

import (
	"go/ast"
	"strconv"
	"strings"
)

type comment struct {
	key   jsonTag
	value jsonTag
}

func parseComments(g *ast.CommentGroup) (c comment) {
	if g == nil {
		return
	}

	for _, comment := range g.List {
		t := comment.Text
		if strings.HasPrefix(t, "key:") {
			tag, err := strconv.Unquote(strings.TrimPrefix(t, "key:"))
			if err != nil {
				panic(err)
			}
			c.key = parseJsonTagInner(tag)
		} else if strings.HasPrefix(t, "value:") {
			tag, err := strconv.Unquote(strings.TrimPrefix(t, "value:"))
			if err != nil {
				panic(err)
			}
			c.value = parseJsonTagInner(tag)
		}
	}
	return
}
