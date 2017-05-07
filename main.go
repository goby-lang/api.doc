package main

import (
	"./parser"
	"fmt"
)

func main() {
	class := parser.ClassFromFile("vm/boolean.go")
	fmt.Println(class.Name)
	fmt.Println(class.Line)
	fmt.Println(class.Comment)
}
