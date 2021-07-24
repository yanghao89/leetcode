package leetcode_Top400

import (
	"fmt"
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) (ans string) {
	if len(strs) == 0 {
		return
	}
	ans = strs[0]
	for _, v := range strs {
		for strings.Index(v, ans) != 0 {
			if len(ans) == 0 {
				return
			}
			ans = ans[0 : len(ans)-1]
		}
	}
	return
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"", "b"}))
}
