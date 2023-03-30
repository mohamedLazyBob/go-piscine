package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if IsLower(string(s[i])) {
			runes[i] = rune(runes[i] - 32)
		}
	}
	return string(runes)
}
