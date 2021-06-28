package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	node1 := utils.CreateTree([]int{1, 2, 3})
	node2 := node1
	fmt.Println(IsSameTree(node1, node2))
}

func IsSameTree(node1, node2 *utils.TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && IsSameTree(node1.Left, node2.Right) && IsSameTree(node1.Right, node2.Left)
}
