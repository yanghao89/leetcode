package leetcode_Top400

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) (ans string) {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			/**
			状态转移方程 dp(i,j) = s[i] == s[j] ^(i - 1 < j +1 || dp(i-1,j+1))
			*/
			dp[i][j] = s[i] == s[j] && (i-1 < j+1 || dp[i-1][j+1])
			if dp[i][j] && i-j+1 > len(ans) {
				ans = s[j : i+1]
			}
		}
	}
	return
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
}
