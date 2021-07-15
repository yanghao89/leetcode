package hashmap

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func numDifferentIntegers(word string) int {
	maps := make(map[string]struct{}, 0)
	for _, v := range strings.FieldsFunc(word, unicode.IsLower) {
		for len(v) > 1 && v[0] == '0' {
			v = v[1:]
		}
		maps[v] = struct{}{}
	}
	return len(maps)
}

func TestNumDifferentIntegers(t *testing.T) {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
	//fmt.Println(numDifferentIntegers("a123bc34d8ef34"))
	//fmt.Println(numDifferentIntegers("leet1234code234"))
	//fmt.Println(numDifferentIntegers("192383183928778851682383a2089984061937879119"))
}
