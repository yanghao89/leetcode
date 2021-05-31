package leetcode_145

import "leetcode/utils"

func PostorderTraversal(root *utils.TreeNode) []int {
	arr := []int{}
	helper(root, &arr)
	return arr
}

func helper(root *utils.TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		helper(root.Left, arr)
	}
	if root.Right != nil {
		helper(root.Right, arr)
	}
	*arr = append(*arr, root.Val)
}
