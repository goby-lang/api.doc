package main

import (
	"./parser"
	// "fmt"
	"io/ioutil"
	"os"
	"strings"
)

func root() string {
	return os.Getenv("GOPATH") + "/src/github.com/rooby-lang/rooby"
}

func dir() string {
	return root() + "/vm"
}

func main() {
	classes := []parser.Class{}
	files, err := ioutil.ReadDir(dir())
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		// fmt.Println(file.Name())
		if strings.Contains(file.Name(), "spec.go") {
			continue
		}
		filename := dir() + "/" + file.Name()
		// fmt.Println("Parsing:", file.Name())
		class := parser.ClassFromFile(filename)
		if class.Line != 0 {
			// fmt.Println("Class found:", class.Name)
			classes = append(classes, class)
		} else {
			// fmt.Println("No class found. Skipped.")
		}
	}

	// fmt.Println(class.Name)
	// fmt.Println(class.Line)
	// fmt.Println(class.Comment)

	parser.Write("./doc.json", classes)
}
