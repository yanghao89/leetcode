package hashmap

import (
	"fmt"
	"math"
	"testing"
)

func findShortestSubArray(nums []int) (ans int) {
	maps := make(map[int][]int, 0)
	dex := 0
	for i, num := range nums {
		maps[num] = append(maps[num], i)
		dex = max(dex, len(maps[num]))
	}
	fmt.Println(maps)
	ans = math.MaxInt64
	for _, item := range maps {
		if len(item) == dex {
			vv := len(item) - 1
			L := item[0]
			dd := item[vv] - L + 1
			ans = int(math.Min(float64(dd), float64(ans)))
		}
	}
	return
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestFindShortestSubArray(t *testing.T) {
	a := "apply"
	b := "apple"
	//fmt.Println(string(a[len(a)-1]), string(b[len(b)-1]))
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i], b[i])
	}
	//fmt.Println(a1, b1)
	//fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
}
