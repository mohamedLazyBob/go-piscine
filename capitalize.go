package piscine

func isLetter(s rune) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z')
}

func isDigit(s rune) bool {
	return s >= '0' && s <= '9'
}

func isAlphanumeric(s rune) bool {
	return isLetter(s) || isDigit(s)
}

func Capitalize(s string) string {
	runes := []rune(s)

	if isLetter(runes[0]) {
		runes[0] = runes[0] - 32
	}

	for i := 1; i < len(runes); i++ {
		if IsUpper(string(runes[i])) {
			runes[i] = runes[i] + 32
		}
		if !isAlphanumeric(runes[i-1]) && IsLower(string(runes[i])) {
			runes[i] = runes[i] - 32
		}
	}
	return string(runes)
}
