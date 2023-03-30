package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 0; a <= 9; a++ {
		for b := a + 1; b <= 9; b++ {
			for c := b + 1; c <= 9; c++ {
				if a != 0 || b != 1 || c != 2 {
					z01.PrintRune(' ')
				}
				z01.PrintRune(rune(48 + a))
				z01.PrintRune(rune(48 + b))
				z01.PrintRune(rune(48 + c))
				if a != 7 || b != 8 || c != 9 {
					z01.PrintRune(',')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
