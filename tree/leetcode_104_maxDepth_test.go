package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestMaxDepth104(t *testing.T) {
	node1 := utils.CreateTree([]int{1, 2, 3, 4, 5})
	//
	fmt.Println(MaxDepth104(node1))
}

//func printNode(root *utils.TreeNode) {
//	if root.Left != nil {
//		printNode(root.Left)
//	}
//	fmt.Println(root.Val)
//	if root.Right != nil {
//		printNode(root.Right)
//	}
//}

func MaxDepth104(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(MaxDepth104(root.Left), MaxDepth104(root.Right)) + 1
}
