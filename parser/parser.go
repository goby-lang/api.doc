package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func ClassFromFile(filename string) Class {
	allMethods := []Method{}
	class := Class{
		Valid: false,
	}

	// Define class name
	qq := strings.Replace(filename, ".go", "", -1)
	p := strings.Replace(qq, "vm/", "", -1)
	class.Name = strings.Title(p)

	// Parse target file
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../rooby/"+filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return class
	}

	// Get comments
	allComments := AllComments{fset, f.Comments}
	// ast.Print(fset, f.Comments)

	// Find class & methods
	var methods *ast.ValueSpec
	// Loop through declarations
	for _, decl := range f.Decls {
		// Continue only for general declarations
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				// Assign class line number if found
				if tSpec, ok := spec.(*ast.TypeSpec); ok && class.MatchName(tSpec.Name.Name) {
					node := tSpec.Name
					class.Line = fset.Position(node.NamePos).Line
					class.Valid = true
				}
				// Assign methods if found
				if vSpec, ok := spec.(*ast.ValueSpec); ok && class.MatchBuiltInMethods(vSpec.Names[0].Name) {
					methods = vSpec
				}
			}
		}
	}

	// Return blank class if class definition is not found
	if !class.Valid {
		return class
	}

	// Retrieve class comments
	class.Comment = allComments.findCommentFor(class.Line)

	// Loop through methods to find each method
	allExpr := methods.Values[0].(*ast.CompositeLit).Elts
	var attrs []ast.Expr
	for _, expr := range allExpr {
		attrs = expr.(*ast.CompositeLit).Elts
		method := Method{}
		// Attributes should only contain "Name" & "Fn" for now
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
