package hashmap

import (
	"fmt"
	"testing"
)

func hasGroupsSizeX(deck []int) (ans bool) {
	maps := make(map[int]int, 0)
	for _, v := range deck {
		maps[v]++
	}
	maxVal := maps[deck[0]]
	ans = true
	for _, v := range maps {
		maxVal = gcd(v, maxVal)
		if maxVal < 2 {
			ans = false
			return
		}
	}
	return
}

func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	}
	return y
}

func testaa(members []int64) {
	if len(members) < 1 {
		fmt.Println(1)
	}
}

func TestHasGroupsSizeX(t *testing.T) {
	ints := [][]int{
		{1, 2, 3, 4, 4, 3, 2, 1},
		{1, 1, 1, 2, 2, 2, 3, 3},
		{1},
		{1, 1},
		{1, 1, 2, 2, 2, 2},
	}
	for i := 0; i < len(ints); i++ {
		fmt.Println(hasGroupsSizeX(ints[i]))
	}
	testaa(nil)
}
