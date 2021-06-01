package leetcode_257

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	node := BinaryTreePaths(root)
	fmt.Println(node)
}
