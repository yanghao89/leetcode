package utils

import (
	"fmt"
	"testing"
)

func TestPreface(t *testing.T) {
	node := CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	//fmt.Println()
	var (
		dfsPre func(node *TreeNode)
		ints   []int
	)
	dfsPre = func(node *TreeNode) {
		if node == nil {
			return
		}
		ints = append(ints, node.Val)
		dfsPre(node.Left)
		dfsPre(node.Right)
	}
	dfsPre(node)
	//前序遍历
	//指对于当前结点，先输出该结点，再输出它的左孩子，最后输出它的右孩子
	fmt.Println(ints)
}

func TestIn(t *testing.T) {
	node := CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	var (
		dfsIn func(root *TreeNode)
		ints  []int
	)
	dfsIn = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfsIn(root.Left)
		ints = append(ints, root.Val)
		dfsIn(root.Right)
	}
	dfsIn(node)
	fmt.Println(ints)
}

func TestAfter(t *testing.T) {
	node := CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	var (
		dfsAfter func(root *TreeNode)
		ints     []int
	)
	dfsAfter = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfsAfter(root.Left)
		dfsAfter(root.Right)
		ints = append(ints, root.Val)
	}
	dfsAfter(node)
	fmt.Println(ints)
}
