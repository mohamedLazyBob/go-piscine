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
		} else if !isLetter(runes[i]) {
			first = true
		}
	}
	return string(runes)
}

func isLetter(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}
