package leetcode_94

import "leetcode/utils"

// InorderTraversal 中序遍历
func InorderTraversal(root *utils.TreeNode) (res []int) {
	var inorder func(root *utils.TreeNode)
	inorder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return
}
