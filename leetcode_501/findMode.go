package leetcode_501

import "leetcode/utils"

func FindMode(root *utils.TreeNode) (ans []int) {
	var (
		dfs                   func(root *utils.TreeNode)
		update                func(x int)
		base, count, maxCount int
	)
	update = func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			ans = append(ans, base)
		} else if count > maxCount {
			maxCount = count
			ans = []int{base}
		}
	}
	dfs = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		update(root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return
}
