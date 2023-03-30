package piscine

import "github.com/01-edu/z01"

func PrintMemory(tab [10]byte) {
	str := ""
	for i, byt := range tab {
		nb := convertHex(int(rune(byt)))
		printWithoutNewLine(nb)
		if ((i+1)%4 == 0 && i != 0) || i == len(tab)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}

		if byt >= 33 && byt <= 126 {
			str += string(rune(byt))
		} else {
			str += "."
		}
	}
	printStr(str)
}

func convertHex(nb int) string {
	numbers := []rune(nil)
	base := "0123456789abcdef"
	for nb != 0 {
		numbers = append(numbers, rune(base[nb%16]))
		nb /= 16
	}
	for i := 0; i < len(numbers)/2; i++ {
		numbers[i], numbers[len(numbers)-i-1] = numbers[len(numbers)-i-1], numbers[i]
	}
	nbHex := string(numbers)
	return nbHex
}

func printWithoutNewLine(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
