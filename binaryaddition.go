package piscine

func BinaryAddition(a int, b int) []int {
	result := a + b
	arr := []int(nil)
	for bit := result; bit > 0; bit >>= 1 {
		arr = append(arr, bit&1)
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
