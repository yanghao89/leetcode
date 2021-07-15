package hashmap

import (
	"fmt"
	"testing"
)

func sumOfUnique(nums []int) (ans int) {
	maps := make(map[int]int)
	for _, v := range nums {
		maps[v]++
	}
	for k, v := range maps {
		if v == 1 {
			ans += k
		}
	}
	return
}

func TestSumOfUnique(t *testing.T) {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 5}))
}
