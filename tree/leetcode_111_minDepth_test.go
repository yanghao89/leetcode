package tree

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

func TestMinDepth(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 9})
	fmt.Println(MinDepth(root))
}

func MinDepth(root *utils.TreeNode) int {
	var (
		dfs func(node *utils.TreeNode) int
	)
	dfs = func(node *utils.TreeNode) int {
		if node == nil {
			return 0
		}
		if (node.Left == nil) && (node.Right == nil) {
			return 1
		}
		min := math.MaxInt16
		if node.Left != nil {
			min = int(math.Min(float64(dfs(node.Left)), float64(min)))
		}
		if node.Right != nil {
			min = int(math.Min(float64(dfs(node.Right)), float64(min)))
		}
		return min + 1
	}
	return dfs(root)
}
