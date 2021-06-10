package leetcode_653

import "leetcode/utils"

func FindTarget(root *utils.TreeNode, k int) bool {
	var (
		check  = make(map[int]bool, 0)
		dfs    func(root *utils.TreeNode, k int)
		isTrue bool
	)
	dfs = func(root *utils.TreeNode, k int) {
		if root == nil {
			return
		}
		dfs(root.Left, k)
		if _, ok := check[k-root.Val]; ok {
			isTrue = ok
			return
		}
		check[root.Val] = true
		dfs(root.Right, k)
	}
	dfs(root, k)
	return isTrue
}
