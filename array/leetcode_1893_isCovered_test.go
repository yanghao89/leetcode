package array

import (
	"fmt"
	"testing"
)

func isCovered(ranges [][]int, left int, right int) bool {
	diff := make([]int, 52, 52)
	for _, v := range ranges {
		diff[v[0]]++
		diff[v[1]+1]--
	}
	cur := 0
	for i := 1; i <= 50; i++ {
		cur += diff[i]
		if i >= left && i <= right && cur <= 0 {
			return false
		}
	}
	return true
}

func TestIsCovered(t *testing.T) {
	fmt.Println(isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
}
