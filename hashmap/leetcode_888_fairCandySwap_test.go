package hashmap

import (
	"fmt"
	"testing"
)

func fairCandySwap(aliceSizes []int, bobSizes []int) (ans []int) {
	maps := make(map[int]struct{}, 0)
	sumA := 0
	for _, v := range aliceSizes {
		sumA += v
		maps[v] = struct{}{}
	}
	sumB := 0
	for _, v := range bobSizes {
		sumB += v
	}
	total := (sumA - sumB) / 2
	for i := 0; ; i++ {
		y := bobSizes[i]
		x := y + total
		if _, ok := maps[x]; ok {
			return []int{x, y}
		}
	}
}

func TestFairCandySwap(t *testing.T) {
	arr := [][]int{
		[]int{1, 1},
		[]int{2, 2},
		[]int{1, 2},
		[]int{2, 3},
		[]int{2},
		[]int{1, 3},
		[]int{1, 2, 5},
		[]int{2, 4},
	}
	for i := 0; i < len(arr); i += 2 {
		fmt.Println(fairCandySwap(arr[i], arr[i+1]))
	}
}
