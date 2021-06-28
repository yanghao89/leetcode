package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func PrintLeftNode(root *utils.TreeNode) []int {
	var (
		ints []int
		dfs  func(root *utils.TreeNode, depth int)
	)
	dfs = func(node *utils.TreeNode, depth int) {
		if node == nil {
			return
		}
		dfs(node.Left, 1)
		if node.Left == nil && node.Right == nil && depth == 1 {
			ints = append(ints, node.Val)
		}
		dfs(node.Right, 2)
	}
	//
	dfs(root, 0)
	return ints
}

func TestPrintLeftNode(t *testing.T) {
	//ints := []int{}
	node := utils.CreateTree([]int{11, 22, 33, 45, 57, 76, 89, 100, 101, 102})
	fmt.Println(PrintLeftNode(node))
}
