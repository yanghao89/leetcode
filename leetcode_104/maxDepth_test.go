package leetcode_104

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	node1 := utils.CreateTree([]int{1, 2, 3, 4, 5})
	printNode(node1)
	fmt.Println(MaxDepth(node1))
}

func printNode(root *utils.TreeNode) {
	if root.Left != nil {
		printNode(root.Left)
	}
	fmt.Println(root.Val)
	if root.Right != nil {
		printNode(root.Right)
	}
}
