// Write a program that prints the name of the program.
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ss := []rune(os.Args[0])

	for i := range ss {
		if i > 1 {
			z01.PrintRune(ss[i])
		}
	}
	z01.PrintRune('\n')
}
