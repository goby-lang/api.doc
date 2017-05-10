package view

type Classes []Class

type Class struct {
	Methods Methods `json:"methods"`
	Name    string  `json:"name"`
	Line    int     `json:"line"`
	Comment string  `json:"comment"`
}
