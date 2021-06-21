package tree

import "leetcode/utils"

func IsUnivalTree(root *utils.TreeNode) bool {
	var (
		ints = make([]int, 0)
		dfs  func(node *utils.TreeNode)
	)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		ints = append(ints, node.Val)
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	if len(ints) > 0 {
		for _, v := range ints {
			if v != ints[0] {
				return false
			}
		}
	}
	return true
}
