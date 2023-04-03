package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	var res []int
	k := 96
	for _, w := range name {
		if w == "--upper" {
			k = 64
			continue
		}
		num := 0
		for _, j := range w {
			num = num*10 + int(rune(j)-'0')
		}
		res = append(res, num)
	}
	for i := 0; i < len(res); i++ {
		if len(res) == 0 {
			break
		} else if res[i] > 26 {
			z01.PrintRune(rune(' '))
			continue
		}
		z01.PrintRune(rune(res[i]) + rune(k))
	}
	if len(res) > 0 {
		z01.PrintRune('\n')
	}
}
