package piscine

func Divisors(n int) int {
	if n < 0 || n > 99999999 {
		return 0
	}
	count := 2
	for i := 2; i < n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}
