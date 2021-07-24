package leetcode_Top400

import (
	"fmt"
	"testing"
)

func groupAnagrams(strs []string) (ans [][]string) {
	m := make(map[[26]int][]string, 0)
	for _, str := range strs {
		cnt := [26]int{}
		for _, v := range str {
			cn := v - 'a'
			cnt[int(cn)]++
		}
		m[cnt] = append(m[cnt], str)
	}
	for _, v := range m {
		ans = append(ans, append([]string(nil), v...))
	}
	return
}

func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
