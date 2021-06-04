package leetcode_606

import (
	"leetcode/utils"
	"strconv"
)

func Tree2str(root *utils.TreeNode) string {
	return helper(root)
}

func helper(root *utils.TreeNode) string {
	if root == nil {
		return ""
	}
	s := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return s
	}
	s += "(" + helper(root.Left) + ")"
	if root.Right != nil {
		s += "(" + helper(root.Right) + ")"
	}
	return s
}
