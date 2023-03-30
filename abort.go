package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}

	for i := range arr {
		for j := range arr {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[len(arr)/2]
}
