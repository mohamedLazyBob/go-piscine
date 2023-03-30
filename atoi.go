package piscine

func Atoi(s string) int {
	number := 0
	if len(s) > 0 {
		factor := 1
		for i := len(s) - 1; i >= 0; i-- {
			if (s[i] < '0' || s[i] > '9') && (i != 0 || (s[0] != '-' && s[0] != '+')) {
				return 0
			}
			if s[i] != '-' && s[i] != '+' {
				number += (int(s[i]) - 48) * factor
				factor = factor * 10
			}
		}
		if s[0] == '-' {
			return -number
		}
	}
	return number
}
