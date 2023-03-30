package piscine

import "github.com/01-edu/z01"

func PrintRange(start, end int) {
	if start < 0 && end < 0 || start > 9 && end > 9 {
		z01.PrintRune('\n')
		return
	}
	if start < 0 {
		start = 0
	} else if start > 9 {
		start = 9
	}
	if end < 0 {
		end = 0
	} else if end > 9 {
		end = 9
	}
	for i := start; ; {
		z01.PrintRune(rune(i) + 48)
		if i != end {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
			break
		}
		if start > end {
			i--
		} else {
			i++
		}
	}
}
