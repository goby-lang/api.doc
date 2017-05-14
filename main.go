package main

import (
	"./parser"
	"./view"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Settings struct {
	GobyPath string `yaml:"gobypath"`
}

func GOPATH() string {
	if os.Getenv("GOPATH") == "" {
		panic("Environment varialbe 'GOPATH' is not set. Setup before continue.")
	}
	return os.Getenv("GOPATH")
}

func GetSettings() Settings {
	settings := Settings{}
	yamlFile, err := ioutil.ReadFile("settings.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &settings)
	if err != nil {
		panic(err)
	}
	return settings
}

func root() string {
	return GOPATH() + GetSettings().GobyPath
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
