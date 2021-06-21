package tree

import "leetcode/utils"

func isSymmetric(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	var (
		dfs func(node1 *utils.TreeNode, node2 *utils.TreeNode) bool
	)
	dfs = func(node1 *utils.TreeNode, node2 *utils.TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) || (node1.Val != node2.Val) {
			return false
		}
		return dfs(node1.Left, node1.Right) && dfs(node2.Left, node2.Right)
	}
	return dfs(root.Left, root.Right)
}
