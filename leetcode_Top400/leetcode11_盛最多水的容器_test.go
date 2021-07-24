package leetcode_Top400

import (
	"fmt"
	"testing"
)

func maxArea(height []int) (ans int) {
	L, R := 0, len(height)-1
	for L < R {
		are := min(height[L], height[R]) * (R - L)
		ans = max(ans, are)
		//小于等于 指针>> | <<指针
		if height[L] <= height[R] {
			L++
		} else {
			R--
		}
	}
	return
}

func TestMaxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
