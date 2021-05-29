package leetcode_108

import "leetcode/utils"

func SortedArrayToBST(arr []int) *utils.TreeNode {
	return helper(arr, 0, len(arr)-1)
}

func helper(arr []int, left, right int) *utils.TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)>>1
	root := &utils.TreeNode{Val: arr[mid]}
	root.Left = helper(arr, left, mid-1)
	root.Right = helper(arr, mid+1, right)
	return root
}
