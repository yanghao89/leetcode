package medium

import (
	"fmt"
	"testing"
)

func findRepeatedDnaSequences(s string) (ans []string) {
	maps := make(map[string]bool, 0)
	for i, n := 0, len(s)-9; i < n; i++ {
		if _, ok := maps[s[i:i+10]]; ok {
			maps[s[i:i+10]] = ok
		} else {
			maps[s[i:i+10]] = false
		}
	}
	for k, v := range maps {
		if v {
			ans = append(ans, k)
		}
	}
	return
}

func TestFindRepeatedDnaSequences(t *testing.T) {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
