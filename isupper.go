package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] < 'A' || runes[i] > 'Z' {
			return false
		}
	}
	return true
}
