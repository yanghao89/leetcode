package hashmap

import (
	"fmt"
	"testing"
)

func repeatedNTimes(nums []int) (ans int) {
	maps := make(map[int]int, 0)
	for _, v := range nums {
		maps[v]++
	}
	for k, v := range maps {
		if v == len(nums)/2 {
			ans = k
			return
		}
	}
	return
}

func TestRepeatedNTimes(t *testing.T) {
	ints := [][]int{
		{1, 2, 3, 3},
		{2, 1, 2, 5, 3, 2},
		{5, 1, 5, 2, 5, 3, 5, 4},
	}
	for i := 0; i < len(ints); i++ {
		fmt.Println(repeatedNTimes(ints[i]))
	}

}
