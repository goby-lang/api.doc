package parser

type Methods []Method

type Method struct {
	FnName  string `json:"name"`
	FnLine  int    `json:"line"`
	Comment string `json:"desc"`
}
