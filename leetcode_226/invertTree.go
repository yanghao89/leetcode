package leetcode_226

import "leetcode/utils"

func InvertTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	Left := InvertTree(root.Left)
	right := InvertTree(root.Right)
	root.Left = right
	root.Right = Left
	return root
}
