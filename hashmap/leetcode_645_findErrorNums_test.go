package hashmap

import (
	"fmt"
	"testing"
)

func findErrorNums(nums []int) (ans []int) {
	maps := make(map[int]int, 0)
	for _, num := range nums {
		maps[num]++
	}
	ans = make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if c := maps[i]; c == 2 {
			ans[0] = i
		} else if c == 0 {
			ans[1] = i
		}
	}
	return
}

func TestFindErrorNums(t *testing.T) {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}
