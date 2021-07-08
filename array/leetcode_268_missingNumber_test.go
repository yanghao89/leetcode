package array

import (
	"fmt"
	"testing"
)

func missingNumber(nums []int) int {
	res := len(nums)
	for k := range nums {
		res ^= k ^ nums[k]
	}
	return res
}

type STW struct {
	A *string `json:"a"`
}

func TestMissingNumber(t *testing.T) {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
