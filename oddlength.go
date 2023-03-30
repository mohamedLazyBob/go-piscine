package piscine

func Oddlength(strings []string) []string {
	oddString := []string(nil)
	for _, word := range strings {
		if len(word)%2 != 0 {
			oddString = append(oddString, word)
		}
	}
	return oddString
}
