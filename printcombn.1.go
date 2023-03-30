package piscine

import "github.com/01-edu/z01"

func PrintCombN1(n int) {
	digits := "0123456789"
	k := 0
	start := 0
	limit := Binomial(n)
	actual := rune(start + 48)
	number := ""
	prefix := ""
	if n > 1 {
		prefix = digits[start : start+n-1]
		actual = rune(prefix[len(prefix)-1]) + 1
	}
	for {
		accumulator := prefix + string(actual)
		number += accumulator
		k++
		if k == limit {
			break
		} else {
			actual++
			if actual > '9' && n > 1 && start < 8 {
				start++
				prefix = digits[start : start+n-1]
				actual = rune(prefix[len(prefix)-1]) + 1
			}
			number += ", "
		}
	}
	printStr(number)
}

func Binomial(k int) int {
	if k > 10/2 {
		k = 10 - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (10 - k + i) * b / i
	}
	return b
}

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
