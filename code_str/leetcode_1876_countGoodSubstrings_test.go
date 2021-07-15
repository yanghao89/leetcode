package code_str

import (
	"fmt"
	"testing"
)

func countGoodSubstrings(s string) (ans int) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			ans++
		}
	}
	return
}

func TestCountGoodSubstrings(t *testing.T) {
	fmt.Println(countGoodSubstrings("xyzzaz"))
	fmt.Println(countGoodSubstrings("aababcabc"))
}
