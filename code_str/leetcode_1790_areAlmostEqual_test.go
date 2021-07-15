package code_str

import (
	"fmt"
	"testing"
)

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	var (
		cnt int
	)
	var (
		c1, c2 byte
	)
	for i := 0; i < len(s1); i++ {

		if s1[i] == s2[i] {
			continue
		}
		if cnt == 0 {
			c1 = s1[i]
			c2 = s2[i]
		} else if cnt == 1 {
			if c1 != s2[i] || c2 != s1[i] {
				return false
			}
		} else if cnt == 2 {
			return false
		}
		cnt++
	}
	return cnt != 1
}

func TestAreAlmostEqual(t *testing.T) {
	fmt.Println(areAlmostEqual("bank", "kanb"))
	fmt.Println(areAlmostEqual("attack", "defend"))
}
