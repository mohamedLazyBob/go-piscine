package piscine

func AlphaPosition(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c) - 97 + 1
	} else if c >= 'A' && c <= 'Z' {
		return int(c) - 65 + 1
	} else {
		return -1
	}
}
