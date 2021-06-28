package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func levelOrderBottom(root *utils.TreeNode) [][]int {

	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var (
		dfs func(node *utils.TreeNode, n int)
	)
	dfs = func(node *utils.TreeNode, n int) {
		if node == nil {
			return
		}
		if len(res) == n {
			res = append(res, []int{})
		}
		dfs(node.Left, n+1)
		dfs(node.Right, n+1)
		res[n] = append(res[n], node.Val)
	}
	dfs(root, 0)
	l := len(res)
	for i := 0; i < l/2; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}
	//maps := [][]int{}
	//for i := len(res); i > 0; i-- {
	//	maps = append(maps, res[i-1])
	//}
	return res
}

func TestLevelOrderBottom(t *testing.T) {
	fmt.Println(LevelOrder(utils.CreateTree([]int{3, 7, 9, 15, 20})))
}
