package tree

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

func maxPathSum(root *utils.TreeNode) (ans int) {
	ans = math.MinInt32
	var (
		dfs func(node *utils.TreeNode) int
	)
	dfs = func(node *utils.TreeNode) int {
		if node == nil {
			return 0
		}
		L := max(dfs(node.Left), 0)
		R := max(dfs(node.Right), 0)
		path := node.Val + L + R
		ans = max(ans, path)
		return node.Val + max(L, R)
	}
	dfs(root)
	return
}

func TestMaxPathSum(t *testing.T) {
	fmt.Println(maxPathSum(utils.CreateTree([]int{1, 2, 3})))
}
