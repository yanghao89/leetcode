package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func majorityElement02(nums []int) (ans int) {
	sort.Ints(nums)
	maps := make(map[int]int, 0)
	for _, v := range nums {
		if maps[v] >= len(nums)/2 {
			return v
		}
		maps[v]++
	}
	return -1
}

func TestMajorityElement02(t *testing.T) {
	fmt.Println(majorityElement02([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
