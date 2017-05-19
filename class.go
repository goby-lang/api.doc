package main

import (
	"html/template"
	"strings"
)

type Classes []Class

type Class struct {
	ClassMethods    Methods       `json:"class_methods"`
	InstanceMethods Methods       `json:"instance_methods"`
	Name            string        `json:"name"`
	Line            int           `json:"line"`
	Comment         template.HTML `json:"comment"`
	Filename        string
	Commit          string
	Repo            string
}

func (a *Class) MatchName(str string) bool {
	return a.Name == str || (strings.Contains(str, a.Name) && strings.Contains(str, "Object"))
}

func (a *Class) MatchBuiltInMethods(str string) bool {
	return strings.Contains(str, "builtin") && strings.Contains(str, "Methods")
}
