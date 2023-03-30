package piscine

func AddIfPositive(a int, b int) int {
	if a >= 0 && b >= 0 {
		return a + b
	}
	return 0
}
