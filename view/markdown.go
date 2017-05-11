package view

import (
	"fmt"
	"os"
	"text/template"
)

func generateIndexMD(classes Classes) {
	indexFile, err := os.OpenFile("./docs/index.md", os.O_CREATE|os.O_WRONLY, 0777)
	panicIf(err)
	indexTemplate, err := template.New("index").ParseFiles(
		"./view/templates/md/index.md",
	)
	panicIf(err)
	variables := map[string]interface{}{
		"classes": classes,
		"class":   nil,
	}
	err = indexTemplate.ExecuteTemplate(indexFile, "index", variables)
	panicIf(err)
	fmt.Println("Generated: ./docs/index.md")
}

func generateClassMD(classes Classes, class Class) {
	classFile, err := os.OpenFile("./docs/classes/"+class.Filename+".md", os.O_CREATE|os.O_WRONLY, 0777)
	panicIf(err)
	classTemplate, err := template.New(class.Filename).ParseFiles(
		"./view/templates/md/class.md",
	)
	panicIf(err)
	variables := map[string]interface{}{
		"classes": classes,
		"class":   class,
	}
	err = classTemplate.ExecuteTemplate(classFile, "class", variables)
	panicIf(err)
	fmt.Println("Generated: ./docs/classes/" + class.Filename + ".md")
}

func GenerateMarkdown(classes Classes) {
	generateIndexMD(classes)
	for _, class := range classes {
		generateClassMD(classes, class)
	}
}
