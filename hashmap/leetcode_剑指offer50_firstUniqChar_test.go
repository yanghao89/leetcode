package hashmap

import (
	"fmt"
	"testing"
)

func firstUniqChar(s string) byte {
	maps := make(map[byte]int, 0)
	for _, v := range s {
		maps[byte(v)-'a']++
	}
	for k, v := range s {
		if maps[byte(v)-'a'] == 1 {
			return s[k]
		}
	}
	return ' '
}

func TestFirstUniqChar(t *testing.T) {
	fmt.Println(string(firstUniqChar("abaccdeff")))
}
