package tree

import (
	"leetcode/utils"
	"math"
)

func isValidBST(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	return DfsBst(root, math.MinInt64, math.MaxInt64)
}

func DfsBst(node *utils.TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if min >= node.Val || max <= node.Val {
		return false
	}
	return DfsBst(node.Left, min, node.Val) && DfsBst(node.Right, node.Val, max)
}
