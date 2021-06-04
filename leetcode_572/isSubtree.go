package leetcode_572

import "leetcode/utils"

//是否为镜像树
func IsSubtree(root *utils.TreeNode, p *utils.TreeNode) bool {
	if root == nil {
		return false
	}
	return helper(root, p) || IsSubtree(root.Left, p) || IsSubtree(root.Right, p)
}

func helper(p, q *utils.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return helper(p.Left, q.Left) && helper(p.Right, q.Right)
	}
	return false
}
