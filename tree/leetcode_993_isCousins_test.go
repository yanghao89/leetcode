package tree

import "leetcode/utils"

func isCousins(root *utils.TreeNode, x, y int) bool {
	var (
		xParent, yParent *utils.TreeNode
		xDepth, yDepth   int
		xFound, yFound   bool
		dfs              func(node, parent *utils.TreeNode, depth int)
	)
	dfs = func(node, parent *utils.TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
		if xFound && yFound {
			return
		}
		dfs(node.Left, node, depth+1)
		if xFound && yFound {
			return
		}
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)
	return xDepth == yDepth && xParent != yParent
}
