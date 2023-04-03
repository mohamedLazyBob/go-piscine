package main

import "os"

func main() {
	if len(os.Args) == 4 {
		args := os.Args[1:]
		operand1, isValid1 := Atoi(args[0])
		operator := args[1]
		operand2, isValid2 := Atoi(args[2])
		if isValid1 && isValid2 {
			if operator == "+" {
				if operand1 < 2147483647 && operand2 < 2147483647 {
					PrintNbr(operand1 + operand2)
				}
			} else if operator == "-" {
				if operand1 > -2147483648 && operand2 > -2147483648 {
					PrintNbr(operand1 - operand2)
				}
			} else if operator == "*" {
				if operand1 < 2147483647 && operand2 < 2147483647 {
					PrintNbr(operand1 * operand2)
				}
			} else if operator == "/" {
				if operand2 != 0 {
					PrintNbr(operand1 / operand2)
				} else {
					os.Stdout.WriteString("No division by 0\n")
				}
			} else if operator == "%" {
				if operand2 != 0 {
					PrintNbr(operand1 % operand2)
				} else {
					os.Stdout.WriteString("No modulo by 0\n")
				}
			}
		}
	}
}

func Atoi(s string) (int, bool) {
	number := 0
	if len(s) > 0 {
		factor := 1
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] < '0' || s[i] > '9' {
				if i != 0 || (s[0] != '-' && s[0] != '+') {
					return 0, false
				}
			}
			if s[i] != '-' && s[i] != '+' {
				number += (int(s[i]) - 48) * factor
				factor = factor * 10
			}
		}
		if s[0] == '-' {
			return -number, true
		}
	}
	return number, true
}

func PrintNbr(n int) {
	slicedNumber := []int{}
	index := 0
	isMinIntValue := false
	if n < 0 {
		os.Stdout.WriteString("-")
		if n == -9223372036854775808 {
			n = n + 1
			isMinIntValue = true
		}
		n = -1 * n
	} else if n == 0 {
		os.Stdout.WriteString("0")
	}
	for n > 0 {
		slicedNumber = append(slicedNumber, n%10)
		n = n / 10
		index++
	}
	for i := len(slicedNumber) - 1; i >= 0; i-- {
		if isMinIntValue && i == 0 {
			os.Stdout.WriteString("8")
		} else {
			os.Stdout.WriteString(string(rune(48 + slicedNumber[i])))
		}
	}
	os.Stdout.WriteString("\n")
}
