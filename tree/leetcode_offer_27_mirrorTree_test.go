package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func MirrorTree(root *utils.TreeNode) *utils.TreeNode {
	var (
		dfs func(node *utils.TreeNode) *utils.TreeNode
	)
	dfs = func(node *utils.TreeNode) *utils.TreeNode {
		if node == nil {
			return nil
		}
		node.Left = dfs(node.Right)
		node.Right = dfs(node.Left)
		return node
	}
	dfs(root)
	return root
}

func TestMirrorTree(t *testing.T) {
	fmt.Println(MirrorTree(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})))
}
