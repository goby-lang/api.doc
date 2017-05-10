package view

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFrom(filepath string) Classes {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var classes Classes
	err = json.Unmarshal(bytes, &classes)
	if err != nil {
		panic(err)
	}
	for i, class := range classes {
		name := strings.ToLower(class.Name)
		class.Filename = name
		class.Repo = "https://github.com/goby-lang/goby"
		class.Commit = "f32c1fcbfd7e1df021948de1065d342e95ebd03d"
		classes[i] = class
	}

	return classes
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func generateIndexFile(classes Classes) {
	indexFile, err := os.OpenFile("./docs/index.html", os.O_CREATE|os.O_WRONLY, 0777)
	panicIf(err)
	indexTemplate, err := template.New("index").ParseFiles("./view/templates/index.html", "./view/templates/header.html")
	panicIf(err)
	err = indexTemplate.ExecuteTemplate(indexFile, "index", classes)
	panicIf(err)
	fmt.Println("Generated: ./docs/index.html")
}

func generateClassFile(class Class) {
	classFile, err := os.OpenFile("./docs/"+class.Filename+".html", os.O_CREATE|os.O_WRONLY, 0777)
	panicIf(err)
	classTemplate, err := template.New(class.Filename).ParseFiles("./view/templates/class.html", "./view/templates/header.html")
	panicIf(err)
	err = classTemplate.ExecuteTemplate(classFile, "class", class)
	panicIf(err)
	fmt.Println("Generated: ./docs/" + class.Name + ".html")
}

func Generate(classes Classes) {
	generateIndexFile(classes)
	for _, class := range classes {
		generateClassFile(class)
	}
}
