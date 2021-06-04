package leetcode_543

import (
	"leetcode/utils"
	"math"
)

func DiameterOfBinaryTree(root *utils.TreeNode) int {
	max := 0
	helper(root, &max)
	return max
}

func helper(root *utils.TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	a := helper(root.Left, max)
	b := helper(root.Right, max)
	*max = int(math.Max(float64(a+b), float64(*max)))
	return int(math.Max(float64(a), float64(b))) + 1
}
