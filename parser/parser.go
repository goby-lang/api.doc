package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Class struct {
	Methods Methods
	Name    string
	Line    int
	Comment string
	Valid   bool
}

type Methods []Method

type Method struct {
	Name    *ast.KeyValueExpr
	Fn      *ast.KeyValueExpr
	FnName  string
	FnLine  int
	Comment string
}

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
			// result = result + comment.Text + "\n"
			result = append(result, comment.Text)
		}
	}

	return strings.Join(result, "\n")
}

func (a *Class) MatchName(str string) bool {
	return a.Name == str || (strings.Contains(str, a.Name) && strings.Contains(str, "Object"))
}

func (a *Class) MatchBuiltInMethods(str string) bool {
	return strings.Contains(str, "builtin") && strings.Contains(str, "Methods")
}

func ClassFromFile(filename string) Class {
	allMethods := []Method{}
	class := Class{
		Valid: false,
	}

	qq := strings.Replace(filename, ".go", "", -1)
	p := strings.Replace(qq, "vm/", "", -1)
	class.Name = strings.Title(p)
	// fmt.Println(class.Name)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../rooby/"+filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return class
	}

	allComments := AllComments{fset, f.Comments}
	// ast.Print(fset, f.Comments)

	var methods *ast.ValueSpec
	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				// TODO check if class name matches file name
				if tSpec, ok := spec.(*ast.TypeSpec); ok && class.MatchName(tSpec.Name.Name) {
					node := tSpec.Name
					class.Line = fset.Position(node.NamePos).Line
					class.Valid = true
				}
				if vSpec, ok := spec.(*ast.ValueSpec); ok && class.MatchBuiltInMethods(vSpec.Names[0].Name) {
					methods = vSpec
				}
			}
		}
	}

	// Returns blank class if class definition is not found
	if !class.Valid {
		return class
	}

	class.Comment = allComments.findCommentFor(class.Line)
	// fmt.Println(class.Comment)

	allExpr := methods.Values[0].(*ast.CompositeLit).Elts
	var attrs []ast.Expr
	for _, expr := range allExpr {
		attrs = expr.(*ast.CompositeLit).Elts
		method := Method{}
		for _, attr := range attrs {
			thisExpr := attr.(*ast.KeyValueExpr)
			name := thisExpr.Key.(*ast.Ident).Name
			if name == "Name" {
				method.Name = thisExpr
				method.FnName = thisExpr.Value.(*ast.BasicLit).Value
			}
			if name == "Fn" {
				method.Fn = thisExpr
				method.FnLine = fset.Position(method.Fn.Key.(*ast.Ident).NamePos).Line
				method.Comment = allComments.findCommentFor(method.FnLine)
			}
		}
		allMethods = append(allMethods, method)
	}

	class.Methods = allMethods
	return class
}
