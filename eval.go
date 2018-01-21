package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func eval(code string, context *Context) ([]byte, error) {
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

	file.Write(context.SourceFor(code))
	return exec.Command("go", "run", file.Name()).CombinedOutput()
}
