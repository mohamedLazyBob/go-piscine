package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printNumber(number int) {
	c := '0'
	u := '0'
	sign := 1
	if number < 0 {
		z01.PrintRune('-')
		sign = -1
	}
	signedNumber := number * sign
	for i := 0; i < signedNumber/10; i++ {
		c++
	}
	z01.PrintRune(c)
	for i := 0; i < signedNumber%10; i++ {
		u++
	}
	z01.PrintRune(u)
}

func printStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")
	printNumber(points.x)
	printStr(", y = ")
	printNumber(points.y)
	z01.PrintRune('\n')
}
