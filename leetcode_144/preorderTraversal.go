package leetcode_144

import "leetcode/utils"

func PreorderTraversal(root *utils.TreeNode) []int {
	var (
		dfs func(root *utils.TreeNode)
		arr = make([]int, 0)
	)
	dfs = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		arr = append(arr, root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}

	}
	dfs(root)
	return arr
}
func helper(root *utils.TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	//*arr = append(*arr, root.Val)
	if root.Left != nil {
		helper(root.Left, arr)
	}
	//中序
	*arr = append(*arr, root.Val)
	if root.Right != nil {
		helper(root.Right, arr)
	}
	//后序
	//*arr = append(*arr, root.Val)
}
