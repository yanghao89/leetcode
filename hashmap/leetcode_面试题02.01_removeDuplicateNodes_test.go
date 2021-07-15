package hashmap

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func removeDuplicateNodes01(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	maps := make(map[int]struct{}, 0)
	maps[head.Val] = struct{}{}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if _, ok := maps[cur.Val]; !ok {
			maps[cur.Val] = struct{}{}
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}

func removeDuplicateNodes02(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	slow := head
	for slow != nil {
		fast := slow
		for fast.Next != nil {
			if slow.Val == fast.Next.Val {
				fast.Next = fast.Next.Next
			} else {
				fast = fast.Next
			}
		}
		slow = slow.Next
	}
	return head
}

func TestRemoveDuplicateNodes(t *testing.T) {
	head01 := removeDuplicateNodes01(utils.CreateListNode([]int{1, 2, 3, 3, 2, 1}, true))
	for ; head01 != nil; head01 = head01.Next {
		fmt.Println(head01.Val)
	}
	head02 := removeDuplicateNodes01(utils.CreateListNode([]int{1, 1, 1, 1, 2}, true))
	for ; head02 != nil; head02 = head02.Next {
		fmt.Println(head02.Val)
	}
}
