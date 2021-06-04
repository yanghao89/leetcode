package leetcode_617

import "leetcode/utils"

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
