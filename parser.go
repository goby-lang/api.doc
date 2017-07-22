package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/k0kubun/pp"
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
		if class.Line != 0 || class.ClassMethods != nil || class.InstanceMethods != nil {
			class.Filename = strings.Replace(file.Name(), ".go", "", 1)
			classes = append(classes, class)
		}
	}
	return classes
}

func classFromFile(filepath string) Class {
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
	var classMethods *ast.FuncDecl
	var instanceMethods *ast.FuncDecl
	// Loop through declarations
	for _, decl := range f.Decls {
		// Continue only for general declarations
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {

				// Assign class line number if found
				tSpec, ok := spec.(*ast.TypeSpec)
				if ok && class.MatchName(tSpec.Name.Name) {
					node := tSpec.Name
					class.Line = fset.Position(node.NamePos).Line
				}
			}
		}

		// Continue only for func declarations
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			decl := funcDecl.Name.Name
			// Assign instance methods if found
			if class.MatchInstanceMethods(decl) {
				instanceMethods = funcDecl
				pp.Println("MatchInstanceMethods: ", funcDecl.Name.Name)
			}

			// Assign class methods if found
			if class.MatchClassMethods(decl) {
				classMethods = funcDecl
				pp.Println("MatchClassMethods: ", funcDecl)
			}
		}
	}

	// Return blank class if class definition is not found
	// if class.Line == 0 {
	// 	return class
	// }

	// Retrieve class comments
	comments := allComments.findCommentFor(class.Line)
	class.Comment = template.HTML(comments.Description)

	// Loop through instance methods to find each method
	if classMethods != nil {
		class.ClassMethods = retrieveMethodsFromNode(fset, classMethods, allComments)
	}
	if instanceMethods != nil {
		class.InstanceMethods = retrieveMethodsFromNode(fset, instanceMethods, allComments)
	}
	return class
}

func retrieveMethodsFromNode(fset *token.FileSet, funcSpec *ast.FuncDecl, allComments AllComments) []Method {
	methods := []Method{}

	allStmt := funcSpec.Body.List[0]
	stmt := allStmt.(*ast.ReturnStmt).Results[0]
	exprs := stmt.(*ast.CompositeLit).Elts
	for _, expr := range exprs {
		method := Method{}
		//kvs := exprs.Key.(*ast.Ident).Name
		kvs := expr.(*ast.CompositeLit).Elts
		for _, kv := range kvs {
			k := kv.(*ast.KeyValueExpr).Key
			v := kv.(*ast.KeyValueExpr).Value
			name := k.(*ast.Ident).Name
			// Attributes should only contain "Name" & "Fn" for now
			if name == "Name" {
				method.FnName = strings.Replace(v.(*ast.BasicLit).Value, "\"", "", -1)
				method.FnLine = fset.Position(k.(*ast.Ident).NamePos).Line
			}
			if name == "Fn" {
				methodComments := allComments.findCommentFor(method.FnLine)
				method.Params = methodComments.Params
				method.Returns = methodComments.Returns
				method.Comment = template.HTML(methodComments.Description)
			}
			pp.Println(method)
		}
		methods = append(methods, method)
	}
	return methods
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

func InsertLinkToComment(text string, class_name string) string {
	t := text
	puncs := []string{" ", ",", ".", ";", "\n"}
	class_link := " [" + class_name + "](" + class_name + ".html) "
	split_t := strings.Split(t, "```")
	for i, _ := range split_t {
		if i%2 == 1 {
			continue
		}
		for _, punc := range puncs {
			split_t[i] = strings.Replace(split_t[i], " "+class_name+punc, class_link, -1)
		}
	}
	return strings.Join(split_t, "```")
}

func DirectInsertLinkToComment(text string, class_name string) string {
	class_link := " [" + class_name + "](" + class_name + ".html) "
	return strings.Replace(text, class_name, class_link, -1)
}

func insertClassLinksForMethods(methods Methods, classes Classes) Methods {
	// loop methods in a class
	for i, method := range methods {
		text := string(method.Comment)
		// insert link to method comment
		for _, each_class := range classes {
			text = InsertLinkToComment(text, each_class.Name)
		}
		methods[i].Comment = template.HTML(text)

		// insert link to params
		for j, param := range method.Params {
			c := string(param.Class)
			d := string(param.Description)
			for _, each_class := range classes {
				c = DirectInsertLinkToComment(c, each_class.Name)
				d = DirectInsertLinkToComment(d, each_class.Name)
			}
			param.Class = template.HTML(c)
			param.Description = template.HTML(d)
			methods[i].Params[j] = param
		}

		// insert link to returns
		for j, r := range method.Returns {
			c := string(r.Class)
			d := string(r.Description)
			for _, each_class := range classes {
				c = DirectInsertLinkToComment(c, each_class.Name)
				d = DirectInsertLinkToComment(d, each_class.Name)
			}
			r.Class = template.HTML(c)
			r.Description = template.HTML(d)
			methods[i].Returns[j] = r
		}

	}

	return methods
}

func InsertClassLinks(classes Classes) Classes {
	var returned_classes Classes
	// loop classes
	for _, class := range classes {
		text := string(class.Comment)
		// insert link to class comment
		for _, each_class := range classes {
			text = InsertLinkToComment(text, each_class.Name)
		}

		class.InstanceMethods = insertClassLinksForMethods(class.InstanceMethods, classes)
		class.ClassMethods = insertClassLinksForMethods(class.ClassMethods, classes)

		class.Comment = template.HTML(text)
		returned_classes = append(returned_classes, class)
	}

	return returned_classes
}
