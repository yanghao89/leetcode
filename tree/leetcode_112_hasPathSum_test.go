package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(HasPathSum(root, 4))
}

func HasPathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return HasPathSum(root.Left, targetSum) || HasPathSum(root.Right, targetSum)
}
