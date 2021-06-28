package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
	fmt.Println(SumOfLeftLeaves(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})))
}

func SumOfLeftLeaves(root *utils.TreeNode) int {
	var (
		sum int
		dfs func(root *utils.TreeNode, level int)
	)
	dfs = func(node *utils.TreeNode, level int) {
		if node == nil {
			return
		}
		dfs(node.Left, 1)
		if node.Left == nil && node.Right == nil && level == 1 {
			sum += node.Val
		}
		dfs(node.Right, 2)
	}
	dfs(root, 0)
	return sum
}
