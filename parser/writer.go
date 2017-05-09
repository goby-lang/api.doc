package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Write(filepath string, classes []Class) {
	b, err := json.Marshal(classes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Generated:", filepath)
	err = ioutil.WriteFile(filepath, b, 0644)
	if err != nil {
		panic(err)
	}
}
