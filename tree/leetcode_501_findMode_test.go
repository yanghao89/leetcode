package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func FindMode(root *utils.TreeNode) (ans []int) {
	var (
		dfs                   func(root *utils.TreeNode)
		update                func(x int)
		base, count, maxCount int
	)
	update = func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			ans = append(ans, base)
		} else if count > maxCount {
			maxCount = count
			ans = []int{base}
		}
	}
	dfs = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		update(root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return
}

func TestFindMode(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 5, 6, 7, 7, 8})
	arr := FindMode(root)
	fmt.Println(arr)
}
