package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	first := true
	for i := range runes {
		if isLetter(runes[i]) && first {
			if IsLower(string(runes[i])) {
				runes[i] -= 32
			}
			first = false
		} else if IsUpper(string(runes[i])) {
			runes[i] += 32
		} 
		if isAlphanumeric(runes[i]) {
			first = true
		}
	}
	return string(runes)
}

func isLetter(s rune) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z')
}

func isDigit(s rune) bool {
	return s >= '0' && s <= '9'
}

func isAlphanumeric(s rune) bool {
	return isLetter(s) || isDigit(s)
}
