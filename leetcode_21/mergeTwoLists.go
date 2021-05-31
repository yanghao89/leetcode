package leetcode_21

import "leetcode/utils"

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
