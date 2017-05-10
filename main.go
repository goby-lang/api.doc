package main

import (
	"./parser"
	"./view"
	// "fmt"
	"os"
)

func root() string {
	return os.Getenv("GOPATH") + "/src/github.com/goby-lang/goby"
}

func dir() string {
	return root() + "/vm"
}

func main() {
	classes := parser.ClassesFromDir(dir())
	parser.Write("./doc.json", classes)

	data := view.ReadFrom("./doc.json")
	view.Generate(data)
}
