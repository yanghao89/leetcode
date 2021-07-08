package hashmap

import (
	"fmt"
	"testing"
)

func intersect(nums1, nums2 []int) (ints []int) {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	maps := make(map[int]int, 0)
	for _, v := range nums1 {
		maps[v]++
	}
	for _, num := range nums2 {
		if maps[num] > 0 {
			ints = append(ints, num)
			maps[num]--
		}
	}
	return
}

func TestIntersect(t *testing.T) {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
