package hashmap

import (
	"fmt"
	"math"
	"testing"
)

func findLHS(nums []int) (ans int) {
	maps := make(map[int]int, 0)
	for _, num := range nums {
		maps[num]++
	}
	for k, v := range maps {
		if maps[k+1] != 0 {
			ans = int(math.Max(float64(ans), float64(maps[k+1]+v)))
		}
	}
	return
}

func TestFindLHS(t *testing.T) {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(findLHS([]int{1, 2, 3, 4}))
	fmt.Println(findLHS([]int{1, 1, 1, 1}))

}
