package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	arr := PreorderTraversal(root)
	//a := true
	//b := false
	fmt.Println(arr)
}

func PreorderTraversal(root *utils.TreeNode) []int {
	var (
		dfs func(root *utils.TreeNode)
		arr []int
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
