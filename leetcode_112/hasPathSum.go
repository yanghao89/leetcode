package leetcode_112

import "leetcode/utils"

// 路径总和  递归
func HasPathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return HasPathSum(root.Left, targetSum) || HasPathSum(root.Right, targetSum)
}
