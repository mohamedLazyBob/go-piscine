package piscine

import (
	"github.com/01-edu/z01"
)

func isSafePosition(numberQueen, testedPosition int, queenPosition [8]int) bool {
	for i := 0; i < numberQueen; i++ {
		position := queenPosition[i]
		// if the testedPosition is already taken return false
		// you attempt to place a queen on the same column of an another
		if position == testedPosition {
			return false
		}
		// if the testedPosition is already taken return false
		// you attempt to place a queen on the same rise diagonal of an another
		if position == testedPosition-numberQueen+i {
			return false
		}
		// if the testedPosition is already taken return false
		// you attempt to place a queen on the same fall diagonal of an another
		if position == testedPosition+numberQueen-i {
			return false
		}
	}
	return true
}

func solve(numberQueen int, queenPosition [8]int) {
	if numberQueen == 8 {
		printSolution(queenPosition)
	} else {
		for i := 0; i < 8; i++ {
			if isSafePosition(numberQueen, i, queenPosition) {
				queenPosition[numberQueen] = i
				solve(numberQueen+1, queenPosition)
			}
		}
	}
}

func printSolution(queenPosition [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(queenPosition[i] + 49))
	}
	z01.PrintRune(10)
}

func EightQueens() {
	queenPosition := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	solve(0, queenPosition)
}
