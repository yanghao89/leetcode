package array

import (
	"fmt"
	"testing"
)

func numIdenticalPairs(nums []int) (ans int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				ans++
			}
		}
	}
	return
}

func TestNumIdenticalPairs(t *testing.T) {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}
