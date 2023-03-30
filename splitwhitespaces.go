package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []string(nil)
	lastWhiteSpaceIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if lastWhiteSpaceIndex != i {
				arr = append(arr, s[lastWhiteSpaceIndex:i])
			}
			lastWhiteSpaceIndex = i + 1
		} else if i == len(s)-1 {
			arr = append(arr, s[lastWhiteSpaceIndex:i+1])
			lastWhiteSpaceIndex = i + 1
		}
	}
	return arr
}
