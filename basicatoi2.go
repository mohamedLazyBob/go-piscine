package piscine

func BasicAtoi2(s string) int {
	number := 0
	factor := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		number += (int(s[i]) - 48) * factor
		factor = factor * 10
	}
	return number
}
