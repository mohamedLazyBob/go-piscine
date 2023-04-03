package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	counter := 0
	for s := range args {
		counter = s + 1
	}
	for a := 1; a < counter; a++ {
		for b := a + 1; b < counter; b++ {
			if args[a] > args[b] {
				args[a], args[b] = args[b], args[a]
			}
		}
	}
	for b := 1; b <= counter-1; b++ {
		for _, s := range args[b] {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}
