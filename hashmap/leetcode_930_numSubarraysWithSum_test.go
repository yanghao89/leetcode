package hashmap

import (
	"fmt"
	"testing"
)

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	sum := 0
	maps := make(map[int]int, 0)
	for _, num := range nums {
		maps[sum]++
		sum += num
		ans += maps[sum-goal]
	}
	return
}

func TestNumSubarraysWithSum(t *testing.T) {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
}
