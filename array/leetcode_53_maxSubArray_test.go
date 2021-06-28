package array

import (
	"fmt"
	"testing"
)

func maxSubArray(nums []int) int {
	res, dp := nums[0], 0
	for _, v := range nums {
		dp = max(v, dp+v)
		res = max(res, dp)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{1}))
}
