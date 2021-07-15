package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func frequencySort(nums []int) []int {
	maps := make(map[int]int, 0)
	for _, num := range nums {
		maps[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if maps[nums[i]] == maps[nums[j]] {
			return nums[i] > nums[j]
		}
		return maps[nums[i]] < maps[nums[j]]
	})
	return nums
}

func TestFrequencySort(t *testing.T) {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
}
