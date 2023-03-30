package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if ValidateBase(base) {
		number := NbrBase(nbr, base)
		for i := 0; i < len(number); i++ {
			z01.PrintRune(rune(number[i]))
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}

func NbrBase(nbr int, base string) string {
	baseLen := len(base)
	number := ""
	isNegative := false
	if nbr < 0 {
		isNegative = true
	}
	if nbr != 0 {
		for nbr != 0 {
			mod := nbr % baseLen
			if mod < 0 {
				mod = -mod
			}
			number += string(base[mod])
			nbr /= baseLen
		}
	} else {
		number = string(base[0])
	}
	if isNegative {
		number += "-"
	}
	number = StrRev(number)
	return number
}

func ValidateBase(base string) bool {
	baseLen := len(base)
	if baseLen <= 1 {
		return false
	} else if string(base[0]) == "-" || string(base[0]) == "+" {
		return false
	} else {
		occurrence := 0
		for i := 0; i < baseLen-1; i++ {
			occurrence = 0
			for j := 0; j < baseLen; j++ {
				if base[i] == base[j] {
					occurrence += 1
					if occurrence == 2 {
						return false
					}
				}
			}
		}
		return true
	}
}
