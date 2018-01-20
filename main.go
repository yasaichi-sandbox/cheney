package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const cursor = "cheney> "

func main() {
	context := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(cursor)
	for scanner.Scan() {
		code := scanner.Text()

		out, err := eval(code, context)
		if err == nil {
			context = append(context, "use("+code+")")
		}

		fmt.Println(strings.TrimRight(string(out), "\r\n"))
		fmt.Print(cursor)
	}
}