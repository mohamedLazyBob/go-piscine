package piscine

func Split(s, sep string) []string {
	arr := []string(nil)
	lastSepIndex := 0
	sepLen := len(sep)
	for i := 0; i <= len(s)-sepLen; {
		if s[i:i+sepLen] == sep {
			if lastSepIndex != i {
				arr = append(arr, s[lastSepIndex:i])
			}
			lastSepIndex = i + sepLen
			i += sepLen
		} else if i == len(s)-sepLen {
			arr = append(arr, s[lastSepIndex:])
			lastSepIndex = i + sepLen
		}
		i++
	}
	return arr
}
