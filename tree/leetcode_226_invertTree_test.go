package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestInvertTree(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(root)
	node := InvertTree(root)
	fmt.Println(node)
}

func InvertTree(root *utils.TreeNode) *utils.TreeNode {
	var (
		dfs func(node *utils.TreeNode) *utils.TreeNode
	)
	dfs = func(node *utils.TreeNode) *utils.TreeNode {
		if node == nil {
			return nil
		}
		news := &utils.TreeNode{Val: node.Val}
		left, right := dfs(node.Left), dfs(node.Right)
		news.Left = right
		news.Right = left
		return news
	}
	return dfs(root)
}
