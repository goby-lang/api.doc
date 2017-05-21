package main

import (
	"encoding/json"
	"io/ioutil"
)

type Settings struct {
	GobyPath string `yaml:"gobypath"`
	Repo     string `yaml:"repo"`
	Commit   string `yaml:"commit"`
}

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
