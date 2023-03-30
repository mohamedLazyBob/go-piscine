package piscine

func AppendRange(min, max int) []int {
	arr := []int(nil)
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
