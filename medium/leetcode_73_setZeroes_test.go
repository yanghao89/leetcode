package medium

import (
	"fmt"
	"testing"
)

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, matri := range matrix {
		for j, v := range matri {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, matri := range matrix {
		for j := range matri {
			if row[i] || row[j] {
				matri[j] = 0
			}
		}
	}
}

func TestSetZeroes(t *testing.T) {
	a := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(a)
	fmt.Println(a)
}
