package main

import (
	"./parser"
	"fmt"
	"os"
)

func dir() string {
	return os.Getenv("GOPATH") + "/src/github.com/rooby-lang/rooby/vm"
}

func main() {
	class := parser.ClassFromFile(dir() + "/boolean.go")
	fmt.Println(class.Name)
	fmt.Println(class.Line)
	fmt.Println(class.Comment)
}
