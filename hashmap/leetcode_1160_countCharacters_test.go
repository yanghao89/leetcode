package hashmap

import (
	"fmt"
	"testing"
)

func countCharacters(words []string, chars string) (ans int) {
	maps := make(map[string]int, 0)
	for _, v := range chars {
		maps[string(v)]++
	}
	for _, word := range words {
		wordCnt := make(map[string]int, 0)
		//计数
		for i := 0; i < len(word); i++ {
			wordCnt[string(word[i])]++
		}
		isAns := true
		for _, v := range word {
			if maps[string(v)] < wordCnt[string(v)] {
				isAns = false
				break
			}
		}
		if isAns {
			ans += len(word)
		}
	}
	return
}

func TestCountCharacters(t *testing.T) {
	strs := [][]string{
		{"cat", "bt", "hat", "tree"},
		{"atach"},
		{"hello", "world", "leetcode"},
		{"welldonehoneyr"},
	}
	for i := 0; i < len(strs); i += 2 {
		fmt.Println(countCharacters(strs[i], strs[i+1][0]))
	}
}
