package piscine

func CountNegative(numbers []int) int {
	count := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < 0 {
			count++
		}
	}
	return count
}
