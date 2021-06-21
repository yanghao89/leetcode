package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	node := utils.CreateTree([]int{1, 2, 3})
	fmt.Println(InorderTraversal(node))
}

// InorderTraversal 中序遍历
func InorderTraversal(root *utils.TreeNode) (res []int) {
	var inorder func(root *utils.TreeNode)
	inorder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return
}
