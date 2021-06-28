package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func pathSum(root *utils.TreeNode, targetSum int) [][]int {
	path, ans := make([]int, 0), make([][]int, 0)
	var (
		dfs func(node *utils.TreeNode, targetSum int)
	)

	dfs = func(node *utils.TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, targetSum)
	return ans
}

type BBB struct {
	A int `json:"a"`
}

func bbbbs(bbb []*int) []*int {
	bcc := 221
	bbb[6] = &bcc
	return bbb
}

func TestPathSum(t *testing.T) {
	a, b, c, d, e, f, g := 1, 2, 3, 4, 5, 6, 7
	g = 9
	abc := []*int{&a, &b, &c, &d, &e, &f, &g}
	g = 22
	abc = append(abc, &g)
	bbb := bbbbs(abc)
	abc[6] = &g
	fmt.Println(*abc[6], *bbb[6])
	//abc := []int(nil)
	//fmt.Println(unsafe.Sizeof(abc))
	////a := append(&abc, ab...)
	//fmt.Println(*abc[6], pathSum(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 18))
}
