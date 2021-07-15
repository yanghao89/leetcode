package medium

import (
	"fmt"
	"strings"
	"testing"
)

func longestSubstring(s string, k int) (ans int) {
	cnt := make([]int, 26)
	for _, v := range s {
		cnt[v-'a']++
	}
	var split byte
	for i, c := range cnt {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	for _, subStr := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return
}
func TestLongestSubstring(t *testing.T) {
	fmt.Println(longestSubstring("aaabb", 3))
}
