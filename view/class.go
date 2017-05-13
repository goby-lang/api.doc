package view

import (
	"html/template"
)

type Classes []Class

type Class struct {
	Methods  Methods `json:"methods"`
	Name     string  `json:"name"`
	Line     int     `json:"line"`
	Comment  template.HTML  `json:"comment"`
	Filename string
	Commit   string
	Repo     string
}
