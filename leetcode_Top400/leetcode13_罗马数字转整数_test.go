package leetcode_Top400

import (
	"fmt"
	"testing"
)

var sv = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for k := range s {
		val := sv[s[k]]
		if k < n-1 && val < sv[s[k+1]] {
			ans -= val
		} else {
			ans += val
		}
	}
	return
}

func TestRomanToInt(t *testing.T) {
	fmt.Println(romanToInt("III"))
}
