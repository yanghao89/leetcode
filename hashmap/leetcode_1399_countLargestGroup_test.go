package hashmap

import (
	"fmt"
	"testing"
)

func countLargestGroup(n int) (ans int) {

	maps := make(map[int]int, 0)
	maxv := 0
	for i := 1; i <= n; i++ {
		key, i0 := 0, i
		for i0 != 0 {
			key += i0 % 10
			i0 /= 10
		}
		maps[key]++
		maxv = max(maxv, maps[key])
	}
	for _, v := range maps {
		if v > maxv {
			maxv = v
		}
		if v == maxv {
			ans++
		}
	}
	return
}

func TestCountLargestGroup(t *testing.T) {
	fmt.Println(countLargestGroup(13))
}
