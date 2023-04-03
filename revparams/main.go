// Write a program that prints the arguments received in the command line in reverse order.

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
	for i := len(os.Args) - 1; i > 0; i-- {
		printStr(os.Args[i])
		z01.PrintRune('\n')
	}
}
