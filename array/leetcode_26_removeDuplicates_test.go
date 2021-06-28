package array

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	last := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[last] = nums[i]
			last++
		}
	}
	return last
}

func TestRemoveDuplicates(t *testing.T) {
	//fmt.Println(removeDuplicates([]int{1, 1, 2}))
	//fmt.Println(removeDuplicates([]int{-1, 0, 0, 0, 0, 3, 3}))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
