package piscine

func AtoiBase(s string, base string) int {
	baseLen := len(base)
	if ValidateBase(base) {
		number := 0
		factor := 1
		isNegative := false
		if s[0] == '-' {
			s = s[1:]
			isNegative = true
		}
		numberLen := len(s)
		for i := numberLen - 1; i >= 0; i-- {
			number += Index(base, string(s[i])) * factor
			if !isNegative && number < 0 {
				number = -(number + 1)
			}
			factor *= baseLen
		}
		if isNegative {
			number = -number
		}
		return number
	}
	return 0
}
