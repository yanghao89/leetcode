package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func longestWord(works []string) (ans string) {
	sort.Strings(works)
	hash := make(map[string]struct{}, 0)
	for i := 0; i < len(works); i++ {
		vv := works[i][:len(works[i])-1]
		if _, ok := hash[vv]; ok || len(works[i]) == 1 {
			if len(works[i]) > len(ans) {
				ans = works[i]
			}
			hash[works[i]] = struct{}{}
		}
	}
	return
}

func appendSlice(slc []string) (ans []interface{}) {
	slc = append(slc, "2")
	slc[0] = "xxxxxx"
	fmt.Println("append 之后slice", len(slc), cap(slc), slc)
	return
}

type name struct {
}

func TestLongestWord(t *testing.T) {
	slc := make([]string, 0, 10)
	slc = append(slc, "1", "2", "3")
	fmt.Println(len(slc), cap(slc))
	appendSlice(slc)
	fmt.Println(len(slc), cap(slc), slc)
	//fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	//fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}
