package hashmap

import (
	"fmt"
	"testing"
)

func destCity(paths [][]string) (ans string) {
	maps := make(map[string]struct{}, 0)
	for _, v := range paths {
		maps[v[0]] = struct{}{}
	}
	for _, v := range paths {
		if _, ok := maps[v[1]]; !ok {
			ans = v[1]
			return
		}
	}
	return
}

func TestDestCity(t *testing.T) {
	strs := [][][]string{
		{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
		{{"B", "C"}, {"D", "B"}, {"C", "A"}},
	}
	for i := 0; i < len(strs); i++ {
		fmt.Println(destCity(strs[i]))
	}
}
