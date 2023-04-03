// Write a program that prints the name of the program.
package main

import (
	"os"
	"piscine"
)

func main() {
	piscine.PrintStr(os.Args[0])
	piscine.PrintStr("\n")
}
