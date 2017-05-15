package view

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func ReadFrom(filepath string, repo string, commit string) Classes {
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
		class.Repo = repo
		class.Commit = commit
		classes[i] = class
	}

	return classes
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
