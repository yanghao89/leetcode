package hashmap

import (
	"fmt"
	"testing"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	maps := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := maps[nums[i]]; ok {
			if i-maps[nums[i]] <= k {
				return true
			}
		}
		maps[nums[i]] = i
	}
	return false
}

func TestContainsNearbyDuplicate02(t *testing.T) {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
