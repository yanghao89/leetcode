package hashmap

import (
	"fmt"
	"testing"
)

func findLucky(nums []int) (ans int) {
	ans = -1
	maps := make(map[int]int, 0)
	for _, v := range nums {
		maps[v]++
	}
	for k, v := range maps {
		if k == v {
			ans = max(ans, k)
		}
	}
	return
}

func TestFindLucky(t *testing.T) {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))
}
