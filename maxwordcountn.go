package piscine

func MaxWordCountN(text string, n int) map[string]int {
	wordMap := split(text, []rune{' ', ',', '\'', '.'})
	occurMapNMax := make(map[string]int)
	maxValue := 0
	maxKey := ""
	for i := 0; i < n; i++ {
		j := 0
		for key, value := range wordMap {
			j++
			if value > maxValue {
				maxValue = value
				maxKey = key
			} else if j == len(wordMap) {
				occurMapNMax[maxKey] = maxValue
				wordMap[maxKey] = 0
				maxValue = 0
				maxKey = ""
			}
		}
	}
	return occurMapNMax
}

func split(s string, seps []rune) map[string]int {
	occurMap := make(map[string]int)
	lastWhiteSpaceIndex := 0
	for i := 0; i < len(s); i++ {
		isSeparator := false
		for _, sep := range seps {
			if rune(s[i]) == sep {
				isSeparator = true
				break
			}
		}
		if isSeparator {
			if lastWhiteSpaceIndex != i {
				word := s[lastWhiteSpaceIndex:i]
				_, ok := occurMap[word]
				if ok {
					occurMap[word]++
				} else {
					occurMap[word] = 1
				}
			}
			lastWhiteSpaceIndex = i + 1
		} else if i == len(s)-1 {
			word := s[lastWhiteSpaceIndex : i+1]
			_, ok := occurMap[word]
			if ok {
				occurMap[word]++
			} else {
				occurMap[word] = 1
			}
			lastWhiteSpaceIndex = i + 1
		}
	}
	return occurMap
}
