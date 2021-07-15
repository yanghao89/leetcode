package medium

import (
	"fmt"
	"sort"
	"testing"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = min(arr[i], arr[i-1]+1)
	}
	return arr[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	//fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{2, 2, 1, 2, 1}))
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{100, 1, 1000}))
}
