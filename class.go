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

func (c *Class) MatchName(str string) bool {
	return c.Name == str || (strings.Contains(str, c.Name) && strings.Contains(str, "Object"))
}

func (c *Class) MatchInstanceMethods(str string) bool {
	return (strings.Contains(str, "builtin") || strings.Contains(str, "Builtin")) && strings.Contains(str, "Methods")
}

func (c *Class) MatchClassMethods(str string) bool {
	return strings.Contains(str, "builtin") && strings.Contains(str, "Methods")
}

func (c *Class) SetClassname(filepath string) {
	split_path := strings.Split(filepath, "/")
	filename := split_path[len(split_path)-1]
	filename_no_ext := strings.Replace(filename, ".go", "", -1)

	name := ""
	for _, segment := range strings.Split(filename_no_ext, "_") {
		name = name + strings.Title(segment)
	}
	c.Name = name
}
