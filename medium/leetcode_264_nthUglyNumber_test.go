package medium

import (
	"fmt"
	"testing"
)

func nthUglyNumber(n int) int {
	m2, m3, m5 := 0, 0, 0
	dp := make([]int, n, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		a, b, c := dp[m2]*2, dp[m3]*3, dp[m5]*5
		dp[i] = min(a, min(b, c))
		if dp[i] == a {
			m2++
		}
		if dp[i] == b {
			m3++
		}
		if dp[i] == c {
			m5++
		}
	}
	return dp[n-1]
}

func TestNthUglyNumber(t *testing.T) {
	fmt.Println(nthUglyNumber(10))
}
