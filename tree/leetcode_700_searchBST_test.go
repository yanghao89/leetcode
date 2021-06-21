package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

// 搜索二叉树
func SearchBST(root *utils.TreeNode, k int) *utils.TreeNode {
	var (
		res *utils.TreeNode
		dfs func(root *utils.TreeNode, k int)
	)
	dfs = func(root *utils.TreeNode, k int) {
		if root != nil && res == nil {
			if root.Val == k {
				res = root
				return
			}
			dfs(root.Left, k)
			dfs(root.Right, k)
		}
	}
	dfs(root, k)
	return res
}

func TestSearchBST(t *testing.T) {
	data := SearchBST(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7}), 4)

	fmt.Println(data)
}
