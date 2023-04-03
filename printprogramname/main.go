// Write a program that prints the name of the program.
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	// skip ./ at the beggining	of the path
	printStr(os.Args[0][2:])
	printStr("\n")
}
