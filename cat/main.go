package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		input, _ := io.ReadAll(os.Stdin)
		printStr(string(input))
	} else {
		for _, fileName := range args {
			content, err := os.ReadFile(fileName)
			if err != nil {
				printStr("ERROR: open " + fileName + ": no such file or directory")
				z01.PrintRune('\n')
				os.Exit(1)
			}
			printStr(string(content))
		}
	}
}
