// Write a program that prints the arguments received in the command line.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStrln(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		printStrln(os.Args[i])
	}
}
