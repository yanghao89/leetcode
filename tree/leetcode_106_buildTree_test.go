package tree

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func buildTree106(inorder []int, postorder []int) *utils.TreeNode {
	idxMap := make(map[int]int)
	for k, v := range inorder {
		idxMap[v] = k
	}
	var (
		dfs func(L, R int) *utils.TreeNode
	)
	dfs = func(inorderLeft, inorderRight int) *utils.TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &utils.TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := idxMap[val]

		root.Right = dfs(inorderRootIndex+1, inorderRight)
		root.Left = dfs(inorderLeft, inorderRootIndex-1)
		return root
	}
	return dfs(0, len(inorder)-1)
}

func TestBuildTree(t *testing.T) {
	fmt.Println(buildTree106([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
