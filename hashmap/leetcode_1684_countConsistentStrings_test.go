package hashmap

import (
	"fmt"
	"testing"
)

func countConsistentStrings(allowed string, words []string) (ans int) {
	maps := make(map[byte]struct{})
	for _, i := range allowed {
		maps[byte(i)] = struct{}{}
	}
	for i := 0; i < len(words); i++ {
		flag := true
		for _, v := range words[i] {
			if _, ok := maps[byte(v)]; !ok {
				flag = false
			}
		}
		if flag {
			ans++
		}
	}
	return
}

func TestCountConsistentStrings(t *testing.T) {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
}
