package hashmap

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func secondHighest(s string) (ans int) {
	ans = -1
	maps := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if '0' < s[i] && s[i] > '9' {
			continue
		}
		maps[s[i]] = struct{}{}
	}
	if len(maps) <= 1 {
		return
	}
	res := make([]int, 0, len(maps))
	for k := range maps {
		kk, _ := strconv.Atoi(string(k))
		res = append(res, kk)
	}
	sort.Ints(res)
	return res[len(res)-2]
}

func TestSecondHighest(t *testing.T) {
	//fmt.Println(secondHighest("dfa12321afd"))
	//fmt.Println(secondHighest("abc1111"))
	fmt.Println(secondHighest("xyz"))
	fmt.Println(strings.Contains("bank", "kanb"))
}
