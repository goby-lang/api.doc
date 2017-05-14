package main

import (
	"./parser"
	"./view"
	"os"
)

func GOPATH() string {
	if os.Getenv("GOPATH") == "" {
		panic("Environment varialbe 'GOPATH' is not set. Setup before continue.")
	}
	return os.Getenv("GOPATH")
}

func root() string {
	return GOPATH() + "/src/github.com/goby-lang/goby"
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
	view.GenerateHTML(data)
}
