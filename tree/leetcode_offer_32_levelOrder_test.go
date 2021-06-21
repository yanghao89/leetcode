package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func LevelOrderOffer(root *utils.TreeNode) [][]int {
	var (
		dfs func(node *utils.TreeNode, l int)
	)
	level := make([][]int, 0)
	if root == nil {
		return level
	}
	dfs = func(node *utils.TreeNode, l int) {
		if node == nil {
			return
		}
		if l == len(level) {
			level = append(level, []int{})
		}
		level[l] = append(level[l], node.Val)
		if node.Left != nil {
			dfs(node.Left, l+1)
		}
		if node.Right != nil {
			dfs(node.Right, l+1)
		}
	}
	dfs(root, 0)
	return level
}

func TestLevelOrderOffer(t *testing.T) {
	fmt.Println(LevelOrderOffer(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})))
}
