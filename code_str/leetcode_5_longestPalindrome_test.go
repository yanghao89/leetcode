package code_str

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longstr("babad"))
	fmt.Println(longstr("cbbb"))
	fmt.Println(longstr("abcb"))
}

func longstr(s string) string {
	l := len(s)
	start, m, dp := 0, 1, make([][]bool, l)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = true
		for j := 0; j < i; j++ {
			if s[j] == s[i] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true
			} else {
				dp[j][i] = false
			}
			if dp[j][i] {
				curLen := i - j + 1
				if curLen > m {
					m = curLen
					start = j
				}
			}
		}
	}
	return s[start : start+m]
}
