package leetcode_Top400

import (
	"fmt"
	"testing"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lengthOfLongestSubstring(s string) (ans int) {
	n := len(s)
	m := make(map[byte]int, n)
	R := -1
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for R+1 < n && m[s[R+1]] == 0 {
			m[s[R+1]]++
			R++
		}
		ans = max(ans, R-i+1)
	}
	return
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
