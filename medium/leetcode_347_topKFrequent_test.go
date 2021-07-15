package medium

import (
	"fmt"
	"sort"
	"testing"
)

func topKFrequent(nums []int, k int) []int {
	maps := make(map[int]int, 0)
	s := make([]int, 0)
	for _, v := range nums {
		if _, ok := maps[v]; ok {
			maps[v]++
		} else {
			maps[v] = 1
			s = append(s, v)
		}
	}
	sort.Slice(s, func(i, j int) bool {
		return maps[s[i]] > maps[s[j]]
	})
	return s[:k]
}

func TestTopKFrequent(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
