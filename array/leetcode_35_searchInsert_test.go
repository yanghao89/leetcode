package array

import (
	"fmt"
	"testing"
)

func searchInsert(nums []int, target int) int {
	n := len(nums)
	ans, left, right := n, 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func TestSearchInsert(t *testing.T) {

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
