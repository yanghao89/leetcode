package simple

import (
	"fmt"
	"sort"
	"testing"
)

func search(nums []int, target int) (ans int) {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return 0
	}
	right := sort.SearchInts(nums, target+1) - 1
	return right - left + 1
}

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}
