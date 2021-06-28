package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestLowestCommonAncestor235(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	node := LowestCommonAncestor235(root, root, root)
	fmt.Println(node)
}

func LowestCommonAncestor235(root1, root2, root3 *utils.TreeNode) *utils.TreeNode {
	var (
		dfs func(node1, node2, node3 *utils.TreeNode) *utils.TreeNode
	)
	dfs = func(node1, node2, node3 *utils.TreeNode) *utils.TreeNode {
		v1, v2, v3 := node1.Val, node2.Val, node3.Val
		switch {
		case v1 > v3 && v2 > v3:
			return dfs(node1.Left, node2, node3)
		case v1 < v3 && v2 < v3:
			return dfs(node1.Right, node2, node3)
		default:
			return node1
		}
	}
	return dfs(root1, root2, root3)
}
