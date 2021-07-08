package hashmap

import (
	"fmt"
	"strings"
	"testing"
)

func commonChars(words []string) (ans []string) {
	maps := make(map[string]int, 0)
	for _, word := range words[0] {
		maps[string(word)]++
	}
	for j := 1; j < len(words); j++ {
		//统计words 下标为1以后的数组在 哈希表中出现的次数
		for k, v := range maps {
			if strings.Count(words[j], k) <= v {
				maps[k] = strings.Count(words[j], k)
			}
		}
	}
	for k, n := range maps {
		for i := 0; i < n; i++ {
			ans = append(ans, k)
		}
	}
	return
}

func TestCommonChars(t *testing.T) {
	strs := [][]string{
		{"bella", "label", "roller"},
		{"cool", "lock", "cook"},
	}
	for i := 0; i < len(strs); i++ {
		fmt.Println(commonChars(strs[i]))
	}

}
