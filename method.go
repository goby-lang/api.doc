package main

import (
	"html/template"
)

type Methods []Method

type Method struct {
	FnName  string        `json:"name"`
	FnLine  int           `json:"line"`
	Comment template.HTML `json:"desc"`
}
