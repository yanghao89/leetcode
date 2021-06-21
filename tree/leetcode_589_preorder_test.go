package tree

import (
	"leetcode/utils"
	"testing"
)

func Preorder(root *utils.Node) []int {
	if root == nil {
		return []int{}
	}
	list := []int{root.Val}
	for _, v := range root.Children {
		list = append(list, Preorder(v)...)
	}
	return list
}

func TestPreorder(t *testing.T) {

}
