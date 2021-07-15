package medium

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func buildTree105(preorder []int, inorder []int) *utils.TreeNode {
	var (
		dfs func(p []int, i []int) *utils.TreeNode
	)
	dfs = func(p []int, in []int) *utils.TreeNode {
		if len(p) == 0 {
			return nil
		}
		node := &utils.TreeNode{Val: p[0], Left: nil, Right: nil}
		i := 0
		for ; i < len(in); i++ {
			if in[i] == p[0] {
				break
			}
		}
		node.Left = dfs(p[1:len(in[:i])+1], in[:i])
		node.Right = dfs(p[len(in[:i])+1:], in[i+1:])
		return node
	}
	return dfs(preorder, inorder)
}

func TestBuildTree105(t *testing.T) {
	fmt.Println(buildTree105([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
