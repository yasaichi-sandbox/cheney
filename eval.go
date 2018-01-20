package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const header = `
package main

import "github.com/k0kubun/pp"

func use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func main() {
`

func eval(code string, context []string) ([]byte, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file, err := os.Create(filepath.Join(dir, "cheney.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	fmt.Fprintln(buf, header)
	for _, v := range context {
		fmt.Fprintln(buf, v)
	}
	fmt.Fprintln(buf, makeCodePrintResult(code)+"\n}")
	buf.Flush()

	// f, _ := os.Open(file.Name())
	// defer f.Close()
	// b := make([]byte, 1024)
	// f.Read(b)
	// fmt.Print(string(b))

	return exec.Command("go", "run", file.Name()).CombinedOutput()
}

func makeCodePrintResult(code string) string {
	return "pp.Println(" + code + ")"
}
