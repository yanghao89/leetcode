package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func pathSum03(root *utils.TreeNode, targetSum int) int {
	prefixSum := make(map[int]int, 0)
	prefixSum[0] = 1
	var (
		dfs func(node *utils.TreeNode, cur, sum int) int
	)
	dfs = func(node *utils.TreeNode, cur, sum int) int {
		if node == nil {
			return 0
		}
		cur += node.Val
		cnt := 0
		if v, ok := prefixSum[cur-sum]; ok {
			cnt = v
		}
		prefixSum[cur]++
		cnt += dfs(node.Left, cur, sum)
		cnt += dfs(node.Right, cur, sum)
		prefixSum[cur]--
		return cnt
	}
	return dfs(root, 0, targetSum)
}

func TestPathSum03(t *testing.T) {
	fmt.Println(pathSum03(utils.CreateTree([]int{1, 2, 3, 6, 7, 8, 12, 13, 14, 15, 17}), 9))
}
