package piscine

func Itoa(n int) string {
	number := []rune(nil)
	for n != 0 {
		number = append(number, rune((n%10)+48))
		n /= 10
	}
	for i := 0; i < len(number)/2; i++ {
		number[i], number[len(number)-i-1] = number[len(number)-i-1], number[i]
	}
	return string(number)
}
