package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] <= 31 {
			return false
		}
	}
	return true
}
