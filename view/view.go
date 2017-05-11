package view

import (
	"encoding/json"
	"io/ioutil"
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
