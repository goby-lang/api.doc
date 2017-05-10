package view

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
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

	return classes
}

func Generate(classes Classes) {

}
