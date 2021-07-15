package hashmap

import (
	"fmt"
	"testing"
)

func checkIfPangram(sentence string) (ans bool) {
	maps := make(map[byte]struct{}, 0)
	count := 0
	for _, v := range sentence {
		if _, ok := maps[byte(v)]; !ok {
			maps[byte(v)] = struct{}{}
			count++
		}
		if count == 26 {
			ans = true
			return
		}
	}
	return
}

func TestCheckIfPangram(t *testing.T) {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}
