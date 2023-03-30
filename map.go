package piscine

func Map(f func(int) bool, a []int) []bool {
	var arr []bool

	for _, v := range a {
		arr = append(arr, f(v))
	}
	return arr
}
