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

func printNumber(point int) {
	a := '0'
	for i := 0; i < point/10; i++ {
		a++
	}
	z01.PrintRune(a)
	b := '0'
	for i := 0; i < point%10; i++ {
		b++
	}
	z01.PrintRune(b)
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(points.y)
	z01.PrintRune('\n')
}
