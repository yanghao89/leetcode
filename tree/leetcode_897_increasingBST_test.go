package tree

import "leetcode/utils"

func IncreasingBST(root *utils.TreeNode) *utils.TreeNode {
	var (
		vals = make([]int, 0)
		dfs  = func(node *utils.TreeNode) {}
	)
	dfs = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		vals = append(vals, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	tree := &utils.TreeNode{}
	if len(vals) > 0 {
		curNode := tree
		for _, val := range vals {
			curNode.Right = &utils.TreeNode{Val: val}
			curNode = curNode.Right
		}
	}
	return tree.Right
}
