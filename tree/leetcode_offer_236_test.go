package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func LowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if right != nil {
		return right
	}
	return left
}

func TestLowestCommonAncestor(t *testing.T) {
	fmt.Println(LowestCommonAncestor(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), utils.CreateTree([]int{3}), utils.CreateTree([]int{1})))
}
