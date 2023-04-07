package piscine

func StringToIntSlice(s string) []int {
	var sliced []int
	for i := range s {
		sliced = append(sliced, int(s[i]))
	}
	return sliced
}
