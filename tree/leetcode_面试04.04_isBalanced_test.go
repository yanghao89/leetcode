package tree

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

func IsBalanced(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	var (
		dfs func(node *utils.TreeNode) int
	)
	dfs = func(node *utils.TreeNode) int {
		if node == nil {
			return 0
		}
		return int(math.Max(float64(dfs(node.Left)), float64(dfs(node.Right)))) + 1
	}
	return int(math.Abs(float64(dfs(root.Left)-dfs(root.Right)))) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}
func TestIsBalanced(t *testing.T) {
	fmt.Println(IsBalanced(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})))
}
