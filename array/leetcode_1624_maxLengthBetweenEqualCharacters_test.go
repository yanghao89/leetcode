package array

import (
	"fmt"
	"testing"
)

func maxLengthBetweenEqualCharacters(s string) (ans int) {
	ans = -1
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && j-i-1 > ans {
				ans = j - i - 1
			}
		}
	}
	return
}

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	fmt.Println(maxLengthBetweenEqualCharacters("aa"))
	fmt.Println(maxLengthBetweenEqualCharacters("abca"))
}
