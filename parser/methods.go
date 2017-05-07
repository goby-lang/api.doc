package parser

import (
	"go/ast"
)

type Methods []Method

type Method struct {
	Name    *ast.KeyValueExpr
	Fn      *ast.KeyValueExpr
	FnName  string
	FnLine  int
	Comment string
}
