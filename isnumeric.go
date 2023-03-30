package piscine

func IsNumeric(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] < '0' || runes[i] > '9' {
			return false
		}
	}
	return true
}
