package leetcdoe_offer_24

import "leetcode/utils"

// ReverseList 翻转链表
//func ReverseList(head *utils.ListNode) *utils.ListNode {
//	var pre *utils.ListNode
//	cur := head
//	for cur != nil {
//		next := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = next
//	}
//	return pre
//}

func helper(head *utils.ListNode, val int) *utils.ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
