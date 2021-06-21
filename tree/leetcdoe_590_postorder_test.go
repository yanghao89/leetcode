package tree

import "leetcode/utils"

func Postorder(root *utils.Node) []int {
	if root == nil {
		return []int{}
	}
	var list []int
	for _, v := range root.Children {
		list = append(list, Postorder(v)...)
	}
	list = append(list, root.Val)
	return list
}
