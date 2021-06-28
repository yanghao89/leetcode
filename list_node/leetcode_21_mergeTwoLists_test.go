package list_node

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	res := MergeTwoLists(utils.CreateListNode([]int{1, 2, 2, 3, 3, 4, 5}), utils.CreateListNode([]int{6, 7, 8, 9, 10}))
	next := res.Next
	fmt.Println(res.Val)
	for next.Next != nil {
		next = next.Next
		fmt.Println(next.Val)
	}
}

func MergeTwoLists(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var tmp *utils.ListNode
	if l1.Val >= l2.Val {
		tmp = l2
		tmp.Next = MergeTwoLists(l1, l2.Next)
	} else {
		tmp = l1
		tmp.Next = MergeTwoLists(l1.Next, l2)
	}
	return tmp
}
