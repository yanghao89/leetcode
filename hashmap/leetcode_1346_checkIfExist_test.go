package hashmap

import (
	"fmt"
	"testing"
)

func checkIfExist(nums []int) (ans bool) {

	maps := make(map[int]int, 0)
	for i, v := range nums {
		maps[v] = i
	}
	for k := range nums {
		nums[k] <<= 1
	}
	var (
		k int
	)
	for i, v := range nums {
		if k, ans = maps[v]; i != k && ans {
			return
		}
	}
	return
}

func TestCheckIfExist(t *testing.T) {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{7, 1, 14, 11}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
	fmt.Println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
}
