package piscine

func IsLower(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] < 'a' || runes[i] > 'z' {
			return false
		}
	}
	return true
}
