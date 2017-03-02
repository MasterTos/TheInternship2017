package main

import (
	"fmt"
)

var x [8]int

func main() {
	placeQueen(0)
}

func print() {
	fmt.Println("-----------------------")
	for row := 0; row < 8; row++ {
		fmt.Print("|")
		for column := 0; column < 8; column++ {
			if column == x[row] {
				fmt.Print("Q.|")
			} else {
				fmt.Print(" |")
			}
		}
		fmt.Println()
		fmt.Println("-----------------------")
	}
	fmt.Println()
}

func canPlace(row int, column int) bool {
	for r := 0; r < row; r++ {
		if x[r] == row || abs(x[r]-row) == abs(r-column) {
			return false
		}
	}
	return true
}

func placeQueen(n int) {
	if n == 8 {
		print()
	} else {
		for i := 0; i < 8; i++ {
			if canPlace(i, n) {
				x[n] = i
				placeQueen(n + 1)
			}
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
