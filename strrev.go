package piscine

func StrRev(s string) string {
	reverse := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		reverse[i], reverse[len(s)-1-i] = reverse[len(s)-1-i], reverse[i]
	}
	return string(reverse)
}
