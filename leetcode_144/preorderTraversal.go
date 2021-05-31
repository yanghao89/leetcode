package leetcode_144

import "leetcode/utils"

func PreorderTraversal(root *utils.TreeNode) []int {
	arr := []int{}
	helper(root, &arr)
	return arr
}
func helper(root *utils.TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	if root.Left != nil {
		helper(root.Left, arr)
	}
	//中序
	//*arr = append(*arr, root.Val)
	if root.Right != nil {
		helper(root.Right, arr)
	}
	//后序
	//*arr = append(*arr, root.Val)
}
