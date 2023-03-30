package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	slicedNumber := []int{}
	index := 0
	isMinIntValue := false
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			n = n + 1
			isMinIntValue = true
		}
		n = -1 * n
	} else if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		slicedNumber = append(slicedNumber, n%10)
		n = n / 10
		index++
	}
	for i := len(slicedNumber) - 1; i >= 0; i-- {
		if isMinIntValue && i == 0 {
			z01.PrintRune(rune('8'))
		} else {
			z01.PrintRune(rune(48 + slicedNumber[i]))
		}
	}
}
