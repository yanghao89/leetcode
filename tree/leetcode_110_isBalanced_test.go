package tree

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

func TestIsBalanced110(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(IsBalanced110(root))

}

func IsBalanced110(root *utils.TreeNode) bool {
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
		return max(dfs(node.Left), dfs(node.Right)) + 1
	}
	return math.Abs(float64(dfs(root.Left)-dfs(root.Right))) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
