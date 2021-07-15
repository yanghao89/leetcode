package medium

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func buildTree(inorder []int, postorder []int) *utils.TreeNode {
	maps := make(map[int]int, 0)
	for k, v := range inorder {
		maps[v] = k
	}
	var (
		dfs func(L, R int) *utils.TreeNode
	)
	dfs = func(L, R int) *utils.TreeNode {
		if L > R {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		node := &utils.TreeNode{Val: val}
		mid := maps[val]
		node.Right = dfs(mid+1, R)
		node.Left = dfs(L, mid-1)
		return node
	}
	return dfs(0, len(inorder)-1)
}

func TestBuildTree(t *testing.T) {

	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
