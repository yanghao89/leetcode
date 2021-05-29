package leetcode_104

import (
	"leetcode/utils"
	"math"
)

func MaxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(MaxDepth(root.Left)), float64(MaxDepth(root.Right))) + 1)
}
