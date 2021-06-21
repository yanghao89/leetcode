package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

//是否为镜像树
func IsSubtree(root *utils.TreeNode, p *utils.TreeNode) bool {
	if root == nil {
		return false
	}
	var (
		dfs func(node1, node2 *utils.TreeNode) bool
	)
	dfs = func(node1, node2 *utils.TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val == node2.Val {
			return dfs(node1.Left, node2.Left) && dfs(node1.Right, node2.Right)
		}
		return false
	}
	return dfs(root, p) || IsSubtree(root.Left, p) || IsSubtree(root.Right, p)
}

func TestIsSubtree(t *testing.T) {
	fmt.Println(IsSubtree(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), utils.CreateTree([]int{1, 2, 3})))
}
