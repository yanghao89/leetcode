package code_str

import (
	"fmt"
	"testing"
)

func makeEqual(words []string) bool {
	bytes := make([]int, 26, 26)
	for _, word := range words {
		for _, v := range word {
			bytes[v-'a']++
		}
	}
	n := len(words)
	for _, v := range bytes {
		if v%n != 0 {
			return false
		}
	}
	return true
}

func TestMakeEqual(t *testing.T) {
	fmt.Println(makeEqual([]string{"abc", "aabc", "bc"}))
	fmt.Println(1 % 2)
}
