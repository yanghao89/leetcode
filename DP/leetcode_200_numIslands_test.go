package DP

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	a := make([][]byte, 0, 4)
	a = append(a, []byte{'1', '1', '1', '1', '0'})
	a = append(a, []byte{'1', '1', '0', '1', '0'})
	a = append(a, []byte{'1', '1', '0', '0', '0'})
	a = append(a, []byte{'0', '0', '0', '0', '0'})
	fmt.Println(NumIslands(a))
	b := make([][]byte, 0, 4)
	b = append(b, []byte{'1', '1', '0', '0', '0'})
	b = append(b, []byte{'1', '1', '0', '0', '0'})
	b = append(b, []byte{'0', '0', '1', '0', '0'})
	b = append(b, []byte{'0', '0', '0', '1', '1'})
	fmt.Println(NumIslands(b))
}

func NumIslands(arr [][]byte) int {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return 0
	}
	var (
		dfs func(ar [][]byte, H, L int)
	)
	dfs = func(ar [][]byte, H, L int) {
		hn := len(ar)
		cn := len(ar[0])
		if H < 0 || L < 0 || H >= hn || L >= cn || ar[H][L] == '0' {
			return
		}
		ar[H][L] = '0'
		dfs(ar, H+1, L)
		dfs(ar, H-1, L)
		dfs(ar, H, L+1)
		dfs(ar, H, L-1)
	}
	R, C, N := len(arr), len(arr[0]), 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if arr[i][j] == '1' {
				N++
				dfs(arr, i, j)
			}
		}
	}
	return N
}
