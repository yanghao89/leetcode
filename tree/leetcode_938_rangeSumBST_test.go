package tree

import "leetcode/utils"

func RangeSumBST(root *utils.TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return RangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return RangeSumBST(root.Right, low, high)
	}
	return root.Val + RangeSumBST(root.Left, low, high) + RangeSumBST(root.Right, low, high)
}
