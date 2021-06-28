package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func KthLargest(root *utils.TreeNode, k int) int {
	var (
		dfs func(node *utils.TreeNode)
	)
	ints := []int{}
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		ints = append(ints, node.Val)
		dfs(node.Left)
	}
	dfs(root)
	return ints[k-1]
}

func TestKthLargest(t *testing.T) {
	fmt.Println(KthLargest(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 2))
}
