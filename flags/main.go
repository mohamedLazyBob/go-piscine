package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Show() {
	fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func Sort(s string) {
	var a [19999]int
	for _, c := range s {
		a[int(c)]++
	}
	for i, c := range a {
		for c > 0 {
			z01.PrintRune(rune(i))
			c--
		}
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]
	ans := ""
	ok := false
	SortIt := false
	for _, c := range arg {
		ok = true
		if c == "-h" || c == "--help" {
			Show()
			break
		}
		ln := 0
		for j := range c {
			ln = j + 1
		}
		if ln > 0 {
			if c[0] == '-' {
				if ln > 2 && c[2] == 'i' {
					if ln > 8 {
						ans += c[9:]
					}
				}
				if c[1] == 'i' {
					if ln > 3 {
						ans += c[3:]
					}
				}
			} else {
				ans = c + ans
			}
		}
		if c == "-o" || c == "--order" {
			SortIt = true
		}
	}
	if !ok {
		Show()
	}
	if SortIt {
		Sort(ans)
	} else {
		fmt.Println(ans)
	}
}
