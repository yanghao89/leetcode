package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	arr := PostorderTraversal(root)
	//a := true
	//b := false
	fmt.Println(arr)
}

func PostorderTraversal(root *utils.TreeNode) []int {
	var (
		dfs func(node *utils.TreeNode)
		arr []int
	)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		arr = append(arr, node.Val)
	}
	dfs(root)
	return arr

}
