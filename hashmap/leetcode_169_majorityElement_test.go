package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement01(nums []int) int {
	var (
		count, n int
	)
	for _, num := range nums {
		if count == 0 {
			n = num
		}
		if n == num {
			count++
		} else {
			count--
		}
	}
	return n
}

func TestMajorityElement(t *testing.T) {
	fmt.Println(majorityElement01([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
