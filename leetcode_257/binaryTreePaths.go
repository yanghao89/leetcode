package leetcode_257

import (
	"leetcode/utils"
	"strconv"
)

func BinaryTreePaths(root *utils.TreeNode) []string {
	paths := make([]string, 0)
	helper(root, "", &paths)
	return paths
}

func helper(root *utils.TreeNode, path string, paths *[]string) {
	if root != nil {
		str := path
		str += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*paths = append(*paths, str)
		} else {
			str += "->"
			helper(root.Left, str, paths)
			helper(root.Right, str, paths)
		}
	}
}
