package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if i != j {
				z01.PrintRune(rune(i/10 + 48))
				z01.PrintRune(rune(i%10 + 48))
				z01.PrintRune(' ')
				z01.PrintRune(rune(j/10 + 48))
				z01.PrintRune(rune(j%10 + 48))
				if i != 1 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
