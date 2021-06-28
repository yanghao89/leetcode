package tree

import "leetcode/utils"

func sortedArrayToBST(nums []int) *utils.TreeNode {
	var (
		createTree func(nums []int, L, R int) *utils.TreeNode
	)
	createTree = func(nums []int, L, R int) *utils.TreeNode {
		if L > R {
			return nil
		}

		mid := L + (R-L)>>1
		node := &utils.TreeNode{Val: nums[mid]}
		node.Left = createTree(nums, L, mid-1)
		node.Right = createTree(nums, mid+1, R)
		return node
	}
	return createTree(nums, 0, len(nums)-1)
}
