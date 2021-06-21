package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	fmt.Println(MergeTrees(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), utils.CreateTree([]int{10, 11, 12, 13, 14, 15, 16})))
}

func MergeTrees(r1 *utils.TreeNode, r2 *utils.TreeNode) *utils.TreeNode {
	if r1 == nil {
		return r2
	}
	if r2 == nil {
		return r2
	}
	newTree := &utils.TreeNode{Val: r1.Val + r2.Val}
	newTree.Left = MergeTrees(r1.Left, r2.Left)
	newTree.Right = MergeTrees(r1.Right, r2.Right)
	return newTree
}
