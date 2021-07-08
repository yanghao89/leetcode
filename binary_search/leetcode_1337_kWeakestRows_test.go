package binary_search

import (
	"fmt"
	"sort"
	"testing"
)

func kWeakestRows(mat [][]int, k int) []int {
	base := len(mat)
	res := make([]int, base)
	for i, v := range mat {
		res[i] = base*searchRows(v) + i
	}
	sort.Ints(res)
	for i := 0; i < k; i++ {
		res[i] %= base
	}
	return res[:k]
}

func searchRows(rows []int) int {
	L, R := 0, len(rows)-1
	for L <= R {
		min := L + (R-L)>>1
		if rows[min] == 1 {
			L = min + 1
		} else {
			R = min - 1
		}
	}
	return L
}

func TestKWeakestRows(t *testing.T) {
	bbb := make([][]int, 0)
	bbb = append(bbb, []int(nil), []int(nil), []int(nil), []int(nil), []int(nil))
	bbb[0] = []int{1, 1, 0, 0, 0}
	bbb[1] = []int{1, 1, 1, 1, 0}
	bbb[2] = []int{1, 0, 0, 0, 0}
	bbb[3] = []int{1, 1, 0, 0, 0}
	bbb[4] = []int{1, 1, 1, 1, 1}

	fmt.Println(kWeakestRows(bbb, 3))
}
