package medium

import (
	"fmt"
	"testing"
)

func isValidSudoku(board [][]byte) bool {
	rows, columns, boxes := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			tmp := board[i][j] - '1'
			if tmp < 0 || tmp > 8 {
				continue
			}
			k := 3*(i/3) + (j / 3)
			if rows[i][tmp] || columns[tmp][j] || boxes[k][tmp] {
				return false
			}
			rows[i][tmp], columns[tmp][j], boxes[k][tmp] = true, true, true
		}
	}
	return true
}

func TestIsValidSudoku(t *testing.T) {
	var board = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}
