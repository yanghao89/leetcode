package tree

import (
	"fmt"
	"leetcode/utils"
	"strconv"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	node := BinaryTreePaths(root)
	fmt.Println(node)
}

func BinaryTreePaths(root *utils.TreeNode) []string {
	var (
		dfs   func(node *utils.TreeNode, str string)
		paths []string
	)
	dfs = func(node *utils.TreeNode, str string) {
		if node == nil {
			return
		}
		str += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, str)
		} else {
			str += "->"
			dfs(node.Left, str)
			dfs(node.Right, str)
		}
	}
	dfs(root, "")
	return paths
}
