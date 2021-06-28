package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	//
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//
	//left := 0
	//right := len(arr) - 1
	////二分
	//min := left + (right-left)>>1
	//fmt.Println(arr[min])
	node1 := utils.CreateTree([]int{1, 2, 3})
	fmt.Println(IsSymmetric(node1))
}

func IsSymmetric(root *utils.TreeNode) bool {
	var (
		dfs func(node1, node2 *utils.TreeNode) bool
	)
	dfs = func(node1, node2 *utils.TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		return node1.Val == node2.Val && dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}
	return dfs(root, root)
}
