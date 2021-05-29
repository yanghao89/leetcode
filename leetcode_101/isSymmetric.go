package leetcode_101

import "leetcode/utils"

/**
对称二叉树 递归
*/
func IsSymmetric(tree *utils.TreeNode) bool {
	return check(tree, tree)
}

func check(q *utils.TreeNode, p *utils.TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	return q.Val == p.Val && check(q.Left, p.Right) && check(q.Right, p.Left)
}
