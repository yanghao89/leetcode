package leetcode_404

import "leetcode/utils"

func SumOfLeftLeaves(root *utils.TreeNode) int {
	return helper(root)
}
func SumOfLeftLeaves01(root *utils.TreeNode) int {
	ans := 0
	helper01(root, &ans, 0)
	return ans
}

func isNil(root *utils.TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func helper(root *utils.TreeNode) (mid int) {
	if root.Left != nil {
		if isNil(root.Left) {
			mid += root.Left.Val
		} else {
			mid += helper(root.Left)
		}
	}
	if root.Right != nil && !isNil(root.Right) {
		mid += helper(root.Right)
	}
	return mid
}
func helper01(root *utils.TreeNode, sum *int, f int) {
	if root == nil {
		return
	}
	helper01(root.Left, sum, 1)
	if root.Left == nil && root.Right == nil && f == 1 {
		*sum += root.Val
	}
	helper01(root.Right, sum, 2)
}
