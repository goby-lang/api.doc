package view

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

func generateIndexFile(classes Classes) {
	indexFile, err := os.OpenFile("./docs/index.html", os.O_CREATE|os.O_WRONLY, 0777)
	panicIf(err)
	indexTemplate, err := template.New("index").ParseFiles(
		"./view/templates/html/index.html",
		"./view/templates/html/layout.html",
		"./view/templates/html/sidebar.html",
		"./view/templates/html/navbar.html",
	)
	panicIf(err)
	variables := map[string]interface{}{
		"classes": classes,
		"class":   nil,
	}
	err = indexTemplate.ExecuteTemplate(indexFile, "layout", variables)
	panicIf(err)
	fmt.Println("Generated: ./docs/index.html")
}

func generateClassFile(classes Classes, class Class) {
	classFile, err := os.OpenFile("./docs/"+class.Filename+".html", os.O_CREATE|os.O_WRONLY, 0777)
	panicIf(err)
	classTemplate, err := template.New(class.Filename).ParseFiles(
		"./view/templates/html/class.html",
		"./view/templates/html/layout.html",
		"./view/templates/html/sidebar.html",
		"./view/templates/html/navbar.html",
	)
	panicIf(err)
	variables := map[string]interface{}{
		"classes": classes,
		"class":   class,
	}
	err = classTemplate.ExecuteTemplate(classFile, "layout", variables)
	panicIf(err)
	fmt.Println("Generated: ./docs/" + class.Name + ".html")
}

func copyAsset(filename string) {
	bytes, err := ioutil.ReadFile("./view/assets/" + filename)
	panicIf(err)
	_, err = os.OpenFile("./docs/assets/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
	panicIf(err)
	err = ioutil.WriteFile("./docs/assets/"+filename, bytes, 0644)
	panicIf(err)
	fmt.Println("Generated: ./docs/" + filename)
}

func GenerateHTML(classes Classes) {
	os.Mkdir("./docs/assets", 0777)

	generateIndexFile(classes)
	for _, class := range classes {
		generateClassFile(classes, class)
	}
	copyAsset("app.css")
}
