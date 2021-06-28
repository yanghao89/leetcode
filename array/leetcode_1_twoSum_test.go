package array

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int, 0)
	for k, v := range nums {
		if vv, ok := maps[target-v]; ok {
			return []int{vv, k}
		}
		maps[v] = k
	}
	return []int(nil)
}

func TestTwoSum(t *testing.T) {
	ints := make([][]int, 0)
	ints = append(ints, []int(nil))
	ints = append(ints, []int(nil))
	ints = append(ints, []int(nil))
	ints[0] = append(ints[0], 2, 7, 11, 15)
	ints[1] = append(ints[1], 3, 2, 4)
	ints[2] = append(ints[2], 3, 3)
	targetInts := []int{9, 6, 6}
	for i := 0; i < 3; i++ {
		fmt.Println(twoSum(ints[i], targetInts[i]))
	}

}
