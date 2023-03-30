package piscine

func PrintCombN(n int) {
	out := PrintCombNumber('0', n, "")
	printStr(out[:len(out)-2])
}

func PrintCombNumber(start rune, n int, accumulator string) string {
	if n == 0 {
		return accumulator + ", "
	}
	combStr := ""
	for i := start; i <= '9'; i++ {
		digits := accumulator + string(i)
		combStr += PrintCombNumber(i+1, n-1, digits)
	}
	return combStr
}
