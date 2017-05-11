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

	os.RemoveAll("./docs")
	os.Mkdir("./docs", 0777)
	os.Mkdir("./docs/classes", 0777)

	data := view.ReadFrom("./doc.json")
	view.GenerateMarkdown(data)
}
