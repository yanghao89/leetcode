package leetcode_108

import (
	"fmt"
	"leetcode/utils"
	"sync"
	"testing"
)

var (
	wg sync.WaitGroup
)

func TestSortedArrayToBST(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tree := SortedArrayToBST(arr)
	fmt.Println(SortedArrayToBST(arr))
	Printtree(tree)
}

func Printtree(tree *utils.TreeNode) {
	fmt.Println(tree.Val)
	if tree.Left != nil {
		Printtree(tree.Left)
	}
	if tree.Right != nil {
		Printtree(tree.Right)
	}
}
