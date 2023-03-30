package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] < 'a' || runes[i] > 'z' {
			if runes[i] < 'A' || runes[i] > 'Z' {
				if runes[i] < '0' || runes[i] > '9' {
					return false
				}
			}
		}
	}
	return true
}
