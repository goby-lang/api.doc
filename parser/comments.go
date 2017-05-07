package parser

import (
	"go/ast"
	"go/token"
	"strings"
)

type AllComments struct {
	Source   *token.FileSet
	Comments []*ast.CommentGroup
}

func (a *AllComments) findCommentFor(i int) string {
	var comments *ast.CommentGroup
	var result []string
	found := false

	for _, group := range a.Comments {
		for _, comment := range group.List {
			line := a.Source.Position(comment.Slash).Line
			if line == (i - 1) {
				comments = group
				found = true
			}
		}
		if found {
			break
		}
	}

	if found {
		for _, comment := range comments.List {
			result = append(result, comment.Text)
		}
	}

	return strings.Join(result, "\n")
}
