package validation

import (
	"fmt"
	"strconv"
)

func ValidTotal(sudoku [9][9]int) bool {
	for row := range [9]int{} {
		for col := range [9]int{} {
			if !IsValid(row, col, sudoku) {
				return false
			}
		}
	}
	return true
}

func IsValid(row int, col int, sudoku [9][9]int) bool {
	if !validCol(row, col, sudoku) {
		return false
	}
	if !validRow(row, col, sudoku) {
		return false
	}
	if !validateSquare(row, col, sudoku) {
		return false
	}
	return true
}

func validRow(row int, col int, sudoku [9][9]int) bool {
	var num int = sudoku[row][col]
	if num == 0 {
		return true
	}
	for colSudoku := range [9]int{} {
		if sudoku[row][colSudoku] != 0 && num == sudoku[row][colSudoku] && colSudoku != col {
			fmt.Printf("Error al validar por fila\n")
			return false
		}
	}
	return true
}

func validCol(row int, col int, sudoku [9][9]int) bool {
	var num int = sudoku[row][col]
	if num == 0 {
		return true
	}
	for rowSudoku := range [9]int{} {
		if sudoku[rowSudoku][col] != 0 && num == sudoku[rowSudoku][col] && rowSudoku != row {
			fmt.Printf("Error al validar por columna\n")
			return false
		}
	}
	return true
}

func validateSquare(row int, col int, sudoku [9][9]int) bool {
	var rowIndex int = 0
	var colIndex int = 0
	if row >= 6 {
		rowIndex = 6
	} else if row >= 3 {
		rowIndex = 3
	} else if row >= 0 {
		rowIndex = 0
	}

	if col >= 6 {
		colIndex = 6
	} else if col >= 3 {
		colIndex = 3
	} else if col >= 0 {
		colIndex = 0
	}

	var num int = sudoku[row][col]
	for rowAdd := range [3]int{} {
		for colAdd := range [3]int{} {
			if (row != rowIndex+rowAdd && col != colIndex+colAdd) && num == sudoku[rowIndex+rowAdd][colIndex+colAdd] && sudoku[rowIndex+rowAdd][colIndex+colAdd] != 0 {
				fmt.Printf("Numero de error: %s", strconv.Itoa(sudoku[rowIndex+rowAdd][colIndex+colAdd]))
				fmt.Printf("Validando area, row: %s, col %s, rowIndex %s, colIndex %s, rowAdd %s, colAdd %s\n",
					strconv.Itoa(row), strconv.Itoa(col), strconv.Itoa(rowIndex), strconv.Itoa(colIndex), strconv.Itoa(rowAdd), strconv.Itoa(colAdd))
				return false
			}
		}
	}
	return true
}
