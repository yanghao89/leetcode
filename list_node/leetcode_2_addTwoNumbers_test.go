package list_node

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	var (
		cur int
	)
	ret := &utils.ListNode{0, nil}
	p := ret
	for l1 != nil || l2 != nil || cur != 0 {
		var (
			a, b int
		)
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		num := a + b + cur
		cur = num / 10
		p.Next = &utils.ListNode{Val: num % 10, Next: nil}
		p = p.Next
	}
	return ret.Next
}

func TestAddTwoNumbers(t *testing.T) {
	node := addTwoNumbers(utils.CreateListNode([]int{2, 4, 3}), utils.CreateListNode([]int{5, 6, 4}))
	for ; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
