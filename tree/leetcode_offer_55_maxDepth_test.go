package tree

import (
	"leetcode/utils"
	"math"
)

func maxDepth(node *utils.TreeNode) int {
	if node == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(node.Left)), float64(maxDepth(node.Right)))) + 1
}
