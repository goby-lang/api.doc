package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"regexp"
	"strings"
)

func ClassesFromDir(dir string) []Class {
	classes := []Class{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "spec.go") {
			continue
		}
		filename := dir + "/" + file.Name()
		class := classFromFile(filename)
		if class.Line != 0 {
			classes = append(classes, class)
		}
	}
	return classes
}

func classFromFile(filepath string) Class {
	allMethods := []Method{}
	class := Class{}
	class.SetClassname(filepath)

	// Parse target file
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)
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
				}
				// Assign methods if found
				if vSpec, ok := spec.(*ast.ValueSpec); ok && class.MatchInstanceMethods(vSpec.Names[0].Name) {
					methods = vSpec
				}
			}
		}
	}

	// Return blank class if class definition is not found
	if class.Line == 0 {
		return class
	}
	// Retrieve class comments
	class.Comment = template.HTML(allComments.findCommentFor(class.Line))

	// Return class if there is not built-in methods
	if methods == nil {
		return class
	}

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
				method.FnName = strings.Replace(thisExpr.Value.(*ast.BasicLit).Value, "\"", "", -1)
				method.FnLine = fset.Position(thisExpr.Key.(*ast.Ident).NamePos).Line
			}
			if name == "Fn" {
				comments := allComments.findCommentFor(method.FnLine)
				method.Params = ExtractParams(comments)
				method.Returns = ExtractReturns(comments)
				method.Comment = template.HTML(comments)
			}
		}
		allMethods = append(allMethods, method)
	}

	class.InstanceMethods = allMethods
	return class
}

func Write(filepath string, classes []Class) {
	b, err := json.Marshal(classes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Generated:", filepath)
	err = ioutil.WriteFile(filepath, b, 0644)
	if err != nil {
		panic(err)
	}
}

func ExtractParams(comments string) []Param {
	params := []Param{}
	lines := strings.Split(comments, "\n")
	for _, line := range lines {
		matched, err := regexp.MatchString("^ @param", line)
		if err != nil {
			panic(err)
		}
		if matched {
			fmt.Println("MATCHED!!!")
			fmt.Println(line)
			param := Param{}
			words := strings.Split(line, " ")
			words = words[1:len(words)]
			fmt.Println(words)
			if len(words) > 1 {
				fmt.Println(words[1])
				param.Name = words[1]
			}
			if len(words) > 2 {
				fmt.Println(words[2])
				class := words[2]
				class = strings.Replace(class, "[", "", 1)
				class = strings.Replace(class, "]", "", 1)
				param.Class = class
			}
			if len(words) > 3 {
				fmt.Println(words[3:len(words)])
				theRest := strings.Join(words[3:len(words)], " ")
				param.Description = template.HTML(theRest)
			}
			if param.Name != "" {
				params = append(params, param)
			}
		}
	}
	return params
}

func ExtractReturns(comments string) []Return {
	returns := []Return{}
	lines := strings.Split(comments, "\n")
	for _, line := range lines {
		matched, err := regexp.MatchString("^ @return", line)
		if err != nil {
			panic(err)
		}
		if matched {
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
			if r.Class != "" {
				returns = append(returns, r)
			}
		}
	}
	return returns
}
