package tree

import "leetcode/utils"

func sumRootToLeaf(root *utils.TreeNode) int {
	var (
		num int
		dfs func(node *utils.TreeNode, x int)
	)
	dfs = func(node *utils.TreeNode, x int) {
		if node == nil {
			return
		}
		x = x*2 + node.Val
		if node.Left == nil && node.Right == nil {
			num += x
		}
		dfs(node.Left, x)
		dfs(node.Right, x)
	}
	dfs(root, 0)
	return num
}
