package parser

import (
	"go/ast"
	"go/token"
	"strings"
	"github.com/russross/blackfriday"
)

type AllComments struct {
	Source   *token.FileSet
	Comments []*ast.CommentGroup
}

func (a *AllComments) findCommentFor(i int) string {
	var comments *ast.CommentGroup
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
		var result []string
		for _, comment := range comments.List {
			comment.Text = strings.Replace(comment.Text, "// ", "", 1)
			result = append(result, comment.Text)
		}
		joined_result := strings.Join(result, "\n")
		output := blackfriday.MarkdownCommon([]byte(joined_result))
		return string(output)
	}

	return ""
}
