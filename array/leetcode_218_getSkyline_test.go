package array

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type pair struct {
	right  int
	height int
}

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {
	n := len(buildings)
	ints := make([]int, 0, n*2)
	for _, building := range buildings {
		ints = append(ints, building[0], building[1])
	}
	sort.Ints(ints)
	idx := 0
	h := hp{}
	for _, vv := range ints {
		for idx < n && buildings[idx][0] <= vv {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		for len(h) > 0 && h[0].right <= vv {
			heap.Pop(&h)
		}
		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{vv, maxn})
		}
	}
	return
}

func TestGetSkyline(t *testing.T) {
	fmt.Println(getSkyline([][]int{{0, 2, 3}, {2, 5, 3}}))
}
