package tree

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

func MinDiffInBST(root *utils.TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var (
		dfs func(root *utils.TreeNode)
	)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func TestMinDiffInBST(t *testing.T) {
	fmt.Println(MinDiffInBST(utils.CreateTree([]int{1, 2, 3, 4, 6})))
}
