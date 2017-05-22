package main

import (
	"html/template"
)

type Methods []Method

type Method struct {
	FnName  string        `json:"name"`
	FnLine  int           `json:"line"`
	Comment template.HTML `json:"desc"`
	Params  []Param       `json:"params"`
	Returns []Return      `json:"returns"`
}

type Param struct {
	Name        string        `json:"name"`
	Class       template.HTML `json:"class"`
	Description template.HTML `json:"description"`
}

type Return struct {
	Class       template.HTML `json:"string"`
	Description template.HTML `json:"description"`
}
