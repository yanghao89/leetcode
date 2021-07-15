package medium

import (
	"fmt"
	"testing"
)

func majorityElement(nums []int) (ans []int) {
	maps := make(map[int]int, 0)
	for _, v := range nums {
		maps[v]++
	}
	n := len(nums) / 3
	for k, v := range maps {
		if v > n {
			ans = append(ans, k)
		}
	}
	return
}

func TestMajorityElement(t *testing.T) {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
