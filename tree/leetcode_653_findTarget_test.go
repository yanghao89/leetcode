package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func FindTarget(root *utils.TreeNode, k int) bool {
	var (
		check  = make(map[int]bool, 0)
		dfs    func(root *utils.TreeNode, k int)
		isTrue bool
	)
	dfs = func(root *utils.TreeNode, k int) {
		if root == nil {
			return
		}
		dfs(root.Left, k)
		if _, ok := check[k-root.Val]; ok {
			isTrue = ok
			return
		}
		check[root.Val] = true
		dfs(root.Right, k)
	}
	dfs(root, k)
	return isTrue
}

func TestFindTarget(t *testing.T) {
	node := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(FindTarget(node, 9))
}
