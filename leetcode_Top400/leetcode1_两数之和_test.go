package leetcode_Top400

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) (ans []int) {
	m := make(map[int]int, 0)
	for k, v := range nums {
		if val, ok := m[target-v]; ok {
			ans = append(ans, val, k)
			return
		}
		m[v] = k
	}
	return
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
