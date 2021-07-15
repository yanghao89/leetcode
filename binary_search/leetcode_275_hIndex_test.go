package binary_search

import (
	"fmt"
	"sort"
	"testing"
)

func hIndex(citations []int) int {
	n := len(citations)
	L, R := 0, n-1
	for L <= R {
		mid := L + (R-L)>>1
		if citations[mid] >= n-mid {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return n - L
}

func hIndex02(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(i int) bool {
		return citations[i] >= n-i
	})
}

func TestHIndex(t *testing.T) {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
	fmt.Println(hIndex02([]int{0, 1, 3, 5, 6}))
}
