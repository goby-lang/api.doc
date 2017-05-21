package main

import (
	"go/ast"
	"go/token"
	"html/template"
	"regexp"
	"strings"
)

type AllComments struct {
	Source   *token.FileSet
	Comments []*ast.CommentGroup
}

type CommentStruct struct {
	Description string
	Params      []Param
	Returns     []Return
}

func IsParamSpec(line string) bool {
	matched, err := regexp.MatchString("^ @param", line)
	if err != nil {
		panic(err)
	}
	return matched
}

func IsReturnSpec(line string) bool {
	matched, err := regexp.MatchString("^ @return", line)
	if err != nil {
		panic(err)
	}
	return matched
}

func (a *AllComments) findCommentFor(i int) CommentStruct {
	var methodComments CommentStruct
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
			comment.Text = strings.Replace(comment.Text, "//", "", 1)
			if IsParamSpec(comment.Text) {
				param := ExtractParam(comment.Text)
				if param.Name != "" {
					methodComments.Params = append(methodComments.Params, param)
				}
			} else if IsReturnSpec(comment.Text) {
				r := ExtractReturn(comment.Text)
				if r.Class != "" {
					methodComments.Returns = append(methodComments.Returns, r)
				}
			} else {
				result = append(result, comment.Text)
			}
		}
	}
	methodComments.Description = strings.Join(result, "\n")

	return methodComments
}

func ExtractParam(line string) Param {
	param := Param{}
	words := strings.Split(line, " ")
	words = words[1:len(words)]
	if len(words) > 1 {
		param.Name = words[1]
	}
	if len(words) > 2 {
		class := words[2]
		class = strings.Replace(class, "[", "", 1)
		class = strings.Replace(class, "]", "", 1)
		param.Class = class
	}
	if len(words) > 3 {
		theRest := strings.Join(words[3:len(words)], " ")
		param.Description = template.HTML(theRest)
	}
	return param
}

func ExtractReturn(line string) Return {
	r := Return{}
	words := strings.Split(line, " ")
	words = words[1:len(words)]
	if len(words) > 1 {
		class := words[1]
		class = strings.Replace(class, "[", "", 1)
		class = strings.Replace(class, "]", "", 1)
		r.Class = class
	}
	if len(words) > 2 {
		theRest := strings.Join(words[3:len(words)], " ")
		r.Description = template.HTML(theRest)
	}
	return r
}
