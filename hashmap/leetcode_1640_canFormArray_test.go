package hashmap

import (
	"fmt"
	"testing"
)

func canFormArray(arr []int, pieces [][]int) (ans bool) {
	maps := make(map[int]int, 0)
	for k, piece := range pieces {
		maps[piece[0]] = k
	}
	cur := 0
	for cur < len(arr) {
		nums := pieces[maps[arr[cur]]]
		if len(nums) == 0 {
			return
		}
		delete(maps, arr[cur])
		for _, num := range nums {
			if cur >= len(arr) || num != arr[cur] {
				return
			}
			cur++
		}
	}
	return cur == len(arr) && len(maps) == 0
}

func TestCanFormArray(t *testing.T) {
	fmt.Println(canFormArray([]int{85}, [][]int{{85}}))
	fmt.Println(canFormArray([]int{15, 88}, [][]int{{85}, {15}}))
}
