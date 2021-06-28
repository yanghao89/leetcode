package array

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func TestRemoveElement(t *testing.T) {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
