package hashmap

import (
	"fmt"
	"strings"
	"testing"
)

func wordPattern(pattern, s string) bool {
	m1 := make(map[string]byte, 0)
	m2 := make(map[byte]string, 0)
	works := strings.Split(s, " ")
	if len(pattern) != len(works) {
		return false
	}
	for i, work := range works {
		ch := pattern[i]
		if m1[work] > 0 && m1[work] != ch || m2[ch] != "" && m2[ch] != work {
			return false
		}
		m1[work] = ch
		m2[ch] = work
	}
	return true
}

func TestWordPattern(t *testing.T) {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
}
