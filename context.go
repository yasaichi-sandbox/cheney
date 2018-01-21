package main

import (
	"fmt"
	"go/parser"
)

const printerName = "__cheney_p"
const sourceTemplate = `
package main

import (
	"github.com/k0kubun/pp"
)

func ` + printerName + `(val interface{}) {
	pp.Println(val)
}

func main() {
	%s
}
`

type Context struct {
}

func (context *Context) Add(code string) {
	// Do noting for now
}

func (context *Context) SourceFor(code string) string {
	var operation string

	// TODO: Deal with the node types other than ast.Expr
	if isExpr(code) {
		operation = fmt.Sprintf("%s(%s)", printerName, code)
	}

	return fmt.Sprintf(sourceTemplate, operation)
}

func isExpr(code string) bool {
	_, err := parser.ParseExpr(code)
	return err == nil
}
