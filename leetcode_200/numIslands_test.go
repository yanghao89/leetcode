package leetcode_200

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
