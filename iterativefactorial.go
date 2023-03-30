package piscine

func IterativeFactorial(nb int) int {
	if nb > 1 && nb < 2147483647 {
		var result int64 = 1
		for i := nb; i >= 1; i-- {
			result *= int64(i)
		}
		return int(result)
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		return 0
	}
}
