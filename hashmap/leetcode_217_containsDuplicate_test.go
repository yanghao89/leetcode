package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func ContainsDuplicate01(nums []int) bool {
	maps := make(map[int]struct{}, 0)
	for _, v := range nums {
		if _, ok := maps[v]; ok {
			return ok
		}
		maps[v] = struct{}{}
	}
	return false
}

func ContainsDuplicate02(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	fmt.Println(ContainsDuplicate01([]int{1, 2, 3, 1}))
	fmt.Println(ContainsDuplicate02([]int{1, 2, 3, 4}))
}
