package tree

import (
	"leetcode/utils"
	"sync"
)

func sortedListToBST(head *utils.ListNode) *utils.TreeNode {
	listNode := head
	var (
		getNodeLens    func(node *utils.ListNode) int
		createTreeNode func(L, R int) *utils.TreeNode
	)
	getNodeLens = func(node *utils.ListNode) int {
		var num int
		for ; node != nil; node = node.Next {
			num++
		}
		return num
	}
	createTreeNode = func(L, R int) *utils.TreeNode {
		if L > R {
			return nil
		}
		min := L + (R-L)>>1
		node := &utils.TreeNode{}
		node.Left = createTreeNode(L, min-1)
		node.Val = listNode.Val
		listNode = listNode.Next
		node.Right = createTreeNode(min+1, R)
		return node
	}
	return createTreeNode(0, getNodeLens(head)-1)
}

func formatNode(head *utils.ListNode, c chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(c)
	}()
	num := 0
	for ; head != nil; head = head.Next {
		num++
	}
	c <- num
}

func sortedListToBSTChannel(head *utils.ListNode) *utils.TreeNode {
	node := head
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int)
	go formatNode(head, c, &wg)
	num := 0
	for v := range c {
		num = v
	}
	wg.Wait()
	var (
		create func(L, R int) *utils.TreeNode
	)
	create = func(L, R int) *utils.TreeNode {
		if L > R {
			return nil
		}
		mid := L + (R-L)>>1
		root := &utils.TreeNode{}
		root.Left = create(L, mid-1)
		root.Val = node.Val
		node = node.Next
		root.Right = create(mid+1, R)
		return root
	}
	return create(0, num-1)
}
