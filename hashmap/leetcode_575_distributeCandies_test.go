package hashmap

import (
	"fmt"
	"math"
	"testing"
)

func distributeCandies(candyType []int) (ans int) {
	maps := make(map[int]struct{}, 0)
	for _, v := range candyType {
		maps[v] = struct{}{}
	}
	ans = int(math.Min(float64(len(maps)), float64(len(candyType)/2)))
	return
}

func TestDistributeCandies(t *testing.T) {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}
