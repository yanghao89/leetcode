package tree

import (
	"leetcode/utils"
	"testing"
)

func flatten(root *utils.TreeNode) {
	var (
		dfs func(node *utils.TreeNode) []*utils.TreeNode
	)
	dfs = func(node *utils.TreeNode) []*utils.TreeNode {
		list := make([]*utils.TreeNode, 0)
		if node != nil {
			list = append(list, root)
			list = append(list, dfs(node.Left)...)
			list = append(list, dfs(node.Right)...)
		}
		return list
	}
	list := dfs(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
	return
}

func TestFatten(t *testing.T) {
	flatten(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
