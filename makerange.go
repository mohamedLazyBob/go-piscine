package piscine

func MakeRange(min, max int) []int {
	if min < max {
		arr := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			arr[i] = min + i
		}
		return arr
	}
	return []int(nil)
}
