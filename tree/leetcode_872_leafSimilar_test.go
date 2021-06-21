package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func LeafSimilar(root1 *utils.TreeNode, root2 *utils.TreeNode) bool {
	var (
		vals = make([]int, 0)
		dfs  = func(root *utils.TreeNode) {}
	)
	dfs = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		//采用前序遍历,只存储叶子节点的val
		if root.Left == nil && root.Right == nil {
			vals = append(vals, root.Val)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root1)
	val1 := append([]int{}, vals...)
	vals = []int{}
	dfs(root2)
	val2 := append([]int{}, vals...)
	if len(val1) != len(val2) {
		return true
	}
	for k, v := range val1 {
		if v != val2[k] {
			return false
		}
	}
	return true
}

func TestLeafSimilar(t *testing.T) {
	fmt.Println(LeafSimilar(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8}), utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})))
}
