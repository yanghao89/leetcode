package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func relativeSortArray01(arr1 []int, arr2 []int) []int {
	maps := make(map[int]int, 0)
	for i, v := range arr2 {
		maps[v] = i - len(arr2)
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		if r, ok := maps[x]; ok {
			x = r
		}
		if r, ok := maps[y]; ok {
			y = r
		}
		return x < y
	})
	return arr1
}

func relativeSortArray02(arr1, arr2 []int) (ans []int) {
	up := 0
	for _, v := range arr1 {
		if v > up {
			up = v
		}
	}
	f := make([]int, up+1)
	for _, v := range arr1 {
		f[v]++
	}
	ans = make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; f[v] > 0; f[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range f {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return
}

func TestRelativeSortArray(t *testing.T) {
	fmt.Println(relativeSortArray01([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	fmt.Println(relativeSortArray02([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}
