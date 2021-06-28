package tree

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

func rob(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		dfs func(root *utils.TreeNode) []int
	)
	dfs = func(node *utils.TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		l, r := dfs(node.Left), dfs(node.Right)
		selected := node.Val + l[1] + r[1]
		notSelected := int(math.Max(float64(l[0]), float64(l[1]))) + int(math.Max(float64(r[0]), float64(r[1])))
		return []int{selected, notSelected}
	}
	val := dfs(root)
	return int(math.Max(float64(val[0]), float64(val[1])))
}

func TestRob(t *testing.T) {
	fmt.Println(rob(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 9, 10})))
}
