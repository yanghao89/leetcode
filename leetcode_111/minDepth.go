package leetcode_111

import (
	"leetcode/utils"
	"math"
)

func MinDepth(root *utils.TreeNode) int {
	return helper(root)
}

func helper(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	if (root.Left == nil) && (root.Right == nil) {
		return 1
	}
	min := math.MaxInt8
	if root.Left != nil {
		min = int(math.Min(float64(helper(root.Left)), float64(min)))
	}
	if root.Right != nil {
		min = int(math.Min(float64(helper(root.Right)), float64(min)))
	}
	return min + 1
}
