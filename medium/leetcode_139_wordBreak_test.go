package medium

import (
	"fmt"
	"testing"
)

func wordBreak(s string, wordDict []string) bool {
	maps := make(map[string]bool, 0)
	for _, word := range wordDict {
		maps[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && maps[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
