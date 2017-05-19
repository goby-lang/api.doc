package main

import (
	// "./parser"
	// "./view"
	"os"
)

func main() {
	classes := ClassesFromDir(dir())
	Write("./doc.json", classes)

	os.RemoveAll("./docs")
	os.Mkdir("./docs", 0777)
	os.Mkdir("./docs/classes", 0777)

	settings := GetSettings()
	data := ReadFrom("./doc.json", settings.Repo, settings.Commit)
	GenerateHTML(data)
}
