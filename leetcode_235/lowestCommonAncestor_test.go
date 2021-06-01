package leetcode_235

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	node := LowestCommonAncestor(root, root, root)
	fmt.Println(node)
}
