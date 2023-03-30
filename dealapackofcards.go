package piscine

import (
	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	lastCard := 0
	for i := 0; i < 4; i++ {
		printStr("Player " + itoa(i+1) + ": ")
		for i := lastCard; i < lastCard+3; i++ {
			printStr(itoa(deck[i]))
			if i != lastCard+2 {
				printStr(", ")
			}
		}
		z01.PrintRune('\n')
		lastCard += 3
	}
}

func itoa(n int) string {
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
