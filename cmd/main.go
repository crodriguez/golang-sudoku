package main

import (
	"fmt"

	"github.com/crodriguez/solved"
)

func main() {
	var sudoku [9][9]int = [9][9]int{
		{4, 0, 1, 0, 0, 0, 0, 0, 2},
		{0, 0, 5, 2, 0, 7, 3, 0, 0},
		{0, 0, 2, 0, 0, 8, 0, 1, 0},
		{2, 0, 0, 0, 5, 0, 0, 0, 1},
		{5, 0, 0, 3, 0, 1, 0, 0, 9},
		{1, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 5, 0, 9, 0, 0, 1, 0, 0},
		{0, 0, 9, 6, 0, 3, 2, 0, 0},
		{6, 0, 0, 0, 0, 0, 4, 0, 3},
	}
	solved, flag := solved.Solved(sudoku)
	fmt.Print(flag)
	fmt.Print(solved)
}
