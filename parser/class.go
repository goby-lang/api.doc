package parser

import (
	"strings"
)

type Class struct {
	Methods Methods
	Name    string
	Line    int
	Comment string
	Valid   bool
}

func (a *Class) MatchName(str string) bool {
	return a.Name == str || (strings.Contains(str, a.Name) && strings.Contains(str, "Object"))
}

func (a *Class) MatchBuiltInMethods(str string) bool {
	return strings.Contains(str, "builtin") && strings.Contains(str, "Methods")
}
