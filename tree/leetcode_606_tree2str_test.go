package tree

import (
	"fmt"
	"leetcode/utils"
	"strconv"
	"testing"
)

func TestTree2str(t *testing.T) {
	fmt.Println(TreeStr(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))
}

func TreeStr(root *utils.TreeNode) string {
	if root == nil {
		return ""
	}
	s := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return s
	}
	if root.Left != nil {
		s += "(" + TreeStr(root.Left) + ")"
	}
	if root.Right != nil {
		s += "(" + TreeStr(root.Right) + ")"
	}
	return s
}
