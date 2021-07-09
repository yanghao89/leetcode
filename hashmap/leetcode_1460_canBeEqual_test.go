package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func canBeEqualMap(target []int, arr []int) (ans bool) {
	ans = true
	maps := make(map[int]int, 0)
	for _, v := range target {
		maps[v]++
	}
	for _, v := range arr {
		if vm, ok := maps[v]; ok && vm > 0 {
			maps[v]--
		} else {
			ans = false
			return
		}
	}
	return
}
func canBeEquaFor(target []int, arr []int) (ans bool) {
	if len(target) != len(arr) {
		return
	}
	ans = true
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < len(target); i++ {
		if target[i] != arr[i] {
			ans = false
			return
		}
	}
	return
}

func TestCanBeEqual(t *testing.T) {
	fmt.Println(canBeEqualMap([]int{3, 7, 9}, []int{3, 7, 11}))
	fmt.Println(canBeEquaFor([]int{1, 2, 2, 3}, []int{1, 1, 2, 3}))
}
