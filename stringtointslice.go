package piscine

func StringToIntSlice(str string) []int {
	var result []int
	for _, c := range str {
		result = append(result, int(c))
	}
	return result
}
