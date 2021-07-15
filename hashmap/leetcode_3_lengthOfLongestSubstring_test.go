package hashmap

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) (ans int) {

	R, n, m := -1, len(s), make(map[byte]int, 0)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for R+1 < n && m[s[R+1]] == 0 {
			m[s[R+1]]++
			//指针向右移动

			R++
		}
		ans = max(ans, R-i+1)
	}
	return ans
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}
