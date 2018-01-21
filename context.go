package main

import (
	"bytes"
	"fmt"
	"go/parser"
	"regexp"
)

type Context struct {
	packages   []string
	operations []string
}

func NewContext() *Context {
	context := new(Context)
	context.addPackage("github.com/k0kubun/pp")

	return context
}

func (context *Context) Add(code string) {
	importStatement := regexp.MustCompile(`import "([\w/]+)"`)

	if importStatement.MatchString(code) {
		context.addPackage(importStatement.FindStringSubmatch(code)[1])
	} else {
		context.addOperation(code)
	}
}

func (context *Context) SourceFor(code string) []byte {
	buf := bytes.NewBufferString("package main\n")

	fmt.Fprintln(buf, "import (")
	for _, name := range context.packages {
		fmt.Fprintln(buf, `"`+name+`"`)
	}
	fmt.Fprintln(buf, ")")

	fmt.Fprintln(buf, "func main() {")
	for _, op := range context.operations {
		fmt.Fprintln(buf, op)
	}
	if isExpr(code) {
		fmt.Fprintln(buf, "pp.Println("+code+")")
	}
	fmt.Fprintln(buf, "}")

	fmt.Println(string(buf.Bytes()))
	return buf.Bytes()
}

func (context *Context) addPackage(name string) {
	for _, v := range context.packages {
		if name == v {
			return
		}
	}

	context.packages = append(context.packages, name)
}

func (context *Context) addOperation(operation string) {
	if !isExpr(operation) {
		context.operations = append(context.operations, operation)
	}
}

func isExpr(code string) bool {
	_, err := parser.ParseExpr(code)
	return err == nil
}
