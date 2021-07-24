package day

import (
	"fmt"
	"testing"
)

func isCovered(ranges [][]int, left int, right int) (ans bool) {
	ans = true
	dp := make([]int, 52, 52)
	for _, v := range ranges {
		dp[v[0]]++
		dp[v[1]+1]--
	}
	cur := 0
	for i := 1; i <= 50; i++ {
		cur += dp[i]
		// <= left <= right <= 50
		if i >= left && i <= right && cur <= 0 {
			ans = false
			return
		}
	}
	return
}

func TestIsCovered(t *testing.T) {
	fmt.Println(isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
}
