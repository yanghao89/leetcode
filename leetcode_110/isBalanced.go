package leetcode_110

import (
	"leetcode/utils"
	"math"
)

// 平衡二叉树 左边和右边 小于等于1
func IsBalanced(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(helper(root.Left)-helper(root.Right))) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func helper(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(helper(root.Left)), float64(helper(root.Right))) + 1)
}
