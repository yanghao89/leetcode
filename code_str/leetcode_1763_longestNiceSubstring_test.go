package code_str

import (
	"fmt"
	"testing"
)

func longestNiceSubstring(s string) (ans string) {
	for i := 0; i < len(s); i++ {
		a, b := 0, 0
		for j := i; j < len(s); j++ {
			if 'a' <= s[j] && s[j] <= 'z' {
				a |= 1 << (s[j] - 'a')
			} else {
				b |= 1 << (s[j] - 'A')
			}
			if a == b && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return
}

func TestLongestNiceSubstring(t *testing.T) {
	z := byte('z')
	fmt.Println(1 << (z - 'a'))
	fmt.Println(longestNiceSubstring("YazaAay"))
}
