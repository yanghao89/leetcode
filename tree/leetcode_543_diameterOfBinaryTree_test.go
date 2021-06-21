package tree

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

func DiameterOfBinaryTree(root *utils.TreeNode) int {
	max := 0
	diameterOfBinaryTreehelper(root, &max)
	return max
}

func diameterOfBinaryTreehelper(root *utils.TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	a := diameterOfBinaryTreehelper(root.Left, max)
	b := diameterOfBinaryTreehelper(root.Right, max)
	*max = int(math.Max(float64(a+b), float64(*max)))
	return int(math.Max(float64(a), float64(b))) + 1
}

func TestDiameterOfBinaryTree(t *testing.T) {
	fmt.Println(DiameterOfBinaryTree(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})))
}
