package solved

import (
	"fmt"
	"strconv"

	"github.com/crodriguez/validation"
)

func findNextEmpty(sudoku [9][9]int) (int, int) {
	for indexRow := range [9]int{} {
		for indexCol := range [9]int{} {
			if sudoku[indexRow][indexCol] == 0 {
				fmt.Printf("Indices encontrados row: %s, col: %s \n", strconv.Itoa(indexRow), strconv.Itoa(indexCol))
				return indexRow, indexCol
			}
		}
	}
	return -1, -1

}

func Solved(sudoku [9][9]int) ([9][9]int, bool) {
	row, col := findNextEmpty(sudoku)
	if row == -1 && col == -1 {
		return sudoku, true
	}

	for num := 1; num <= 9; num++ {
		fmt.Printf("Num asignado: %s\n", strconv.Itoa(num))
		sudoku[row][col] = num
		if validation.IsValid(row, col, sudoku) {
			sudokuTemp, flag := Solved(sudoku)
			if flag {
				return sudokuTemp, flag
			}
		}
	}
	return sudoku, false
}
