package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

var fns = template.FuncMap{
	"last": func(index int, length int) bool {
		return index == length-1
	},
}

func generateIndexFile(classes Classes) {
	indexFile, err := os.OpenFile("./docs/index.html", os.O_CREATE|os.O_WRONLY, 0777)
	panicIf(err)
	indexTemplate, err := template.New("index").ParseFiles(
		"./templates/html/index.html",
		"./templates/html/layout.html",
		"./templates/html/sidebar.html",
		"./templates/html/navbar.html",
	)
	panicIf(err)
	variables := map[string]interface{}{
		"classes": classes,
		"class":   nil,
		"readme":  template.HTML(readmeHTML("./README.md")),
	}
	err = indexTemplate.ExecuteTemplate(indexFile, "layout", variables)
	panicIf(err)
	fmt.Println("Generated: ./docs/index.html")
}

func generateClassFile(classes Classes, class Class) {
	classFile, err := os.OpenFile("./docs/"+class.Filename+".html", os.O_CREATE|os.O_WRONLY, 0777)
	panicIf(err)
	classTemplate, err := template.New(class.Filename).Funcs(fns).ParseFiles(
		"./templates/html/class.html",
		"./templates/html/layout.html",
		"./templates/html/sidebar.html",
		"./templates/html/navbar.html",
	)
	panicIf(err)
	classComment := blackfriday.MarkdownCommon([]byte(class.Comment))
	class.Comment = template.HTML(string(classComment))
	for i := 0; i < len(class.InstanceMethods); i++ {
		methodComment := blackfriday.MarkdownCommon([]byte(class.InstanceMethods[i].Comment))
		class.InstanceMethods[i].Comment = template.HTML(methodComment)
		for j := 0; j < len(class.InstanceMethods[i].Params); j++ {
			p := class.InstanceMethods[i].Params[j]
			paramClass := blackfriday.MarkdownCommon([]byte(p.Class))
			paramDesc := blackfriday.MarkdownCommon([]byte(p.Description))
			paramClass = stripTag(paramClass)
			paramDesc = stripTag(paramDesc)
			class.InstanceMethods[i].Params[j].Class = template.HTML(paramClass)
			class.InstanceMethods[i].Params[j].Description = template.HTML(paramDesc)
		}
		for j := 0; j < len(class.InstanceMethods[i].Returns); j++ {
			r := class.InstanceMethods[i].Returns[j]
			returnClass := blackfriday.MarkdownCommon([]byte(r.Class))
			returnDesc := blackfriday.MarkdownCommon([]byte(r.Description))
			returnClass = stripTag(returnClass)
			returnDesc = stripTag(returnDesc)
			class.InstanceMethods[i].Returns[j].Class = template.HTML(returnClass)
			class.InstanceMethods[i].Returns[j].Description = template.HTML(returnDesc)
		}
	}
	variables := map[string]interface{}{
		"classes": classes,
		"class":   class,
	}
	err = classTemplate.ExecuteTemplate(classFile, "layout", variables)
	panicIf(err)
	fmt.Println("Generated: ./docs/" + class.Filename + ".html")
}

func copyAsset(filename string) {
	bytes, err := ioutil.ReadFile("./assets/" + filename)
	panicIf(err)
	_, err = os.OpenFile("./docs/assets/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
	panicIf(err)
	err = ioutil.WriteFile("./docs/assets/"+filename, bytes, 0644)
	panicIf(err)
	fmt.Println("Generated: ./docs/" + filename)
}

func readmeHTML(filepath string) string {
	bytes, err := ioutil.ReadFile(filepath)
	panicIf(err)
	html := string(blackfriday.MarkdownCommon(bytes))
	html = strings.Replace(html, "<code class=\"language-", "<code class=\"", -1)
	return html
}

func GenerateHTML(classes Classes) {
	os.Mkdir("./docs/assets", 0777)

	generateIndexFile(classes)
	for _, class := range classes {
		generateClassFile(classes, class)
	}
	copyAsset("app.css")
	copyAsset("app.js")
}

func stripTag(bytes []byte) []byte {
	t := strings.Replace(string(bytes), "<p>", "", 1)
	t = strings.Replace(t, "</p>", "", 1)
	t = strings.Replace(t, "\n", "", 1)
	t = strings.Replace(t, "<a ", "<a class=\"class-link\"", 1)
	return []byte(t)
}
