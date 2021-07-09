package hashmap

import (
	"fmt"
	"testing"
)

func uniqueOccurrences(nums []int) bool {
	m1 := make(map[int]int, 0)
	for _, v := range nums {
		m1[v]++
	}
	m2 := make(map[int]struct{}, 0)
	for _, v := range m1 {
		m2[v] = struct{}{}
	}
	return len(m1) == len(m2)
}

func TestUniqueOccurrences(t *testing.T) {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
