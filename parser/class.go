package parser

import (
	"strings"
)

type Class struct {
	Methods Methods `json:"methods"`
	Name    string  `json:"name"`
	Line    int     `json:"line"`
	Comment string  `json:"comment"`
}

func (a *Class) MatchName(str string) bool {
	return a.Name == str || (strings.Contains(str, a.Name) && strings.Contains(str, "Object"))
}

func (a *Class) MatchBuiltInMethods(str string) bool {
	return strings.Contains(str, "builtin") && strings.Contains(str, "Methods")
}
