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

func toLower(s rune) rune {
	if s >= 'A' && s <= 'Z' {
		return s + 32
	}
	return s
}

func toUpper(s rune) rune {
	if s >= 'a' && s <= 'z' {
		return s - 32
	}
	return s
}

func Capitalize(s string) string {
	runes := []rune(s)

	if IsLower(string(runes[0])) {
		runes[0] = toUpper(runes[0])
	}

	for i := 1; i < len(runes); i++ {
		if IsUpper(string(runes[i])) {
			runes[i] = toLower(runes[i])
		}
		if !isAlphanumeric(runes[i-1]) && IsLower(string(runes[i])) {
			runes[i] = toUpper(runes[i])
		}
	}
	return string(runes)
}
