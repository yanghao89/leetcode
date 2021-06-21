package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func SortedArrayToBST(nums []int) *utils.TreeNode {
	var (
		dfs func(arr []int) *utils.TreeNode
	)
	dfs = func(arr []int) *utils.TreeNode {
		if len(arr) == 0 {
			return nil
		}
		return &utils.TreeNode{
			Val:   arr[len(arr)/2],
			Left:  dfs(arr[:len(arr)/2]),
			Right: dfs(arr[len(arr)/2+1:]),
		}
	}
	return dfs(nums)
}

func TestSortedArrayToBST(t *testing.T) {
	fmt.Println(SortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}
