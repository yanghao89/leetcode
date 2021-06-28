package tree

import "leetcode/utils"

func buildTree(preorder []int, inorder []int) *utils.TreeNode {
	var (
		dfs func(preorder []int, inorder []int) *utils.TreeNode
	)
	dfs = func(preorder []int, inorder []int) *utils.TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		node := &utils.TreeNode{Val: preorder[0], Left: nil, Right: nil}
		i := 0
		for ; i < len(inorder); i++ {
			if inorder[i] == preorder[0] {
				break
			}
		}
		node.Left = dfs(preorder[1:len(inorder[:i])+1], inorder[:i])
		node.Right = dfs(preorder[len(inorder[:i])+1:], inorder[i+1:])
		return node
	}
	return dfs(preorder, inorder)
}
