package hashmap

import (
	"fmt"
	"strings"
	"testing"
)

func uncommonFromSentences(s1 string, s2 string) (ans []string) {
	maps := make(map[string]int, 0)
	for _, v := range strings.Split(s1, " ") {
		maps[v]++
	}
	for _, v := range strings.Split(s2, " ") {
		maps[v]++
	}
	for k, v := range maps {
		if v == 1 {
			ans = append(ans, k)
		}
	}
	return
}

func TestUncommonFromSentences(t *testing.T) {
	strs := []string{"this apple is sweet", "this apple is sour", "apple apple", "banana"}
	for i := 0; i < len(strs); i += 2 {
		fmt.Println(uncommonFromSentences(strs[i], strs[i+1]))
	}
}
