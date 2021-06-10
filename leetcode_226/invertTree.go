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

func dfs(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	L := dfs(root.Left)
	R := dfs(root.Right)
	root.Left = R
	root.Right = L
	return root
}
