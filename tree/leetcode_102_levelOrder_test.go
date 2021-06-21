package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func LevelOrder(root *utils.TreeNode) [][]int {
	var (
		levels = make([][]int, 0)
		dfs    func(node *utils.TreeNode, level int)
	)
	dfs = func(node *utils.TreeNode, level int) {
		if node == nil {
			return
		}
		if len(levels) == level {
			levels = append(levels, []int{})
		}
		levels[level] = append(levels[level], node.Val)
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 0)
	return levels
}

func TestLevelOrder(t *testing.T) {
	fmt.Println(LevelOrder(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))
}
