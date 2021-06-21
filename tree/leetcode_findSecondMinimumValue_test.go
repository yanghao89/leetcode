package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

// 二叉树中第二小的节点
func TestFindSecondMinimumValue(t *testing.T) {
	fmt.Println(FindSecondMinimumValue(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7})))
}
func FindSecondMinimumValue(root *utils.TreeNode) int {
	val1, val2, dfs := -1, root.Val, func(root *utils.TreeNode) {}
	dfs = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		if val1 > 0 && root.Val > val1 {
			return
		}
		if root.Val > val2 {
			val1 = root.Val
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return val1
}
