package leetcode_Top400

import (
	"fmt"
	"testing"
)

var (
	m = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
)

func letterCombinations(s string) (ans []string) {
	if len(s) == 0 {
		return ans
	}
	var dfs func(dig string, idx int, com string)
	dfs = func(dig string, idx int, com string) {
		if idx == len(dig) {
			ans = append(ans, com)
			return
		}
		cc := m[string(dig[idx])]
		for i := 0; i < len(cc); i++ {
			dfs(dig, idx+1, com+string(cc[i]))
		}
	}
	dfs(s, 0, "")
	return
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}
