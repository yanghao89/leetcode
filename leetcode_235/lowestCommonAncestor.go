package leetcode_235

import "leetcode/utils"

func LowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	return helper(root, p, q)
}

func helper(root, p, q *utils.TreeNode) *utils.TreeNode {
	v1, v2, v3 := p.Val, q.Val, root.Val
	switch {
	case v1 > v3 && v2 > v3:
		return helper(root.Right, p, q)
	case v1 < v3 && v2 < v3:
		return helper(root.Left, p, q)
	default:
		return root
	}
}
