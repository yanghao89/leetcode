package medium

import (
	"fmt"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	maps := make(map[[26]int][]string, 0)
	for _, str := range strs {
		cnt := [26]int{}
		for _, v := range str {
			cn := v - 'a'
			cnt[cn]++
		}
		maps[cnt] = append(maps[cnt], str)
	}
	ans := make([][]string, 0, len(maps))
	for _, v := range maps {
		ans = append(ans, v)
	}
	return ans
}

func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
