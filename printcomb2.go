// Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.
// These combinations are separated by a comma and a space.
// without using the rune type, --allow-builtin
// example 00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$

package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if i != 0 || j != 1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(i/10 + 48))
			z01.PrintRune(rune(i%10 + 48))
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10 + 48))
			z01.PrintRune(rune(j%10 + 48))
		}
	}
	z01.PrintRune('\n')
}
