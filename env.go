package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

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
	if settings.Repo == "" {
		panic("Missing 'repo' attribute in /settings.yml")
	}
	if settings.Commit == "" {
		panic("Missing 'commit' attribute in /settings.yml")
	}
	return settings
}

func root() string {
	return GOPATH() + GetSettings().GobyPath
}

func dir() string {
	return root() + "/vm"
}
