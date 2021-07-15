package medium

import (
	"fmt"
	"sort"
	"testing"
)

func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	maps := make(map[int]bool, 0)
	for _, v := range nums {
		maps[v] = true
	}
	long := 0
	for num := range maps {
		if !maps[num-1] {
			cur := num
			start := 1
			for maps[cur+1] {
				cur++
				start++
			}
			if long < start {
				long = start
			}
		}
	}
	return long
}

func TestLongestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
