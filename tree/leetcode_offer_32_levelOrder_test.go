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
	fmt.Println(zigzagLevelOrder(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})))
}

func zigzagLevelOrder(root *utils.TreeNode) [][]int {
	var (
		dfs   func(node *utils.TreeNode, k int)
		level = make([][]int, 0)
	)
	if root == nil {
		return level
	}
	dfs = func(node *utils.TreeNode, k int) {
		if node == nil {
			return
		}
		if len(level) == k {
			level = append(level, []int{})
		}
		level[k] = append(level[k], node.Val)
		if node.Left != nil {
			dfs(node.Left, k+1)
		}
		if node.Right != nil {
			dfs(node.Right, k+1)
		}
	}
	dfs(root, 0)
	for k, v := range level {
		if k%2 > 0 {
			for i, n := 0, len(v); i < n/2; i++ {
				v[i], v[n-i-1] = v[n-i-1], v[i]
			}
		}
	}
	return level
}
