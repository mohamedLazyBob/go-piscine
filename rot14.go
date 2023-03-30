package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, c := range runes {
		if c >= 'a' && c <= 'z' {
			runes[i] = ((c - 97 + 14) % 26) + 97
		} else if c >= 'A' && c <= 'Z' {
			runes[i] = ((c - 65 + 14) % 26) + 65
		}
	}
	return string(runes)
}
