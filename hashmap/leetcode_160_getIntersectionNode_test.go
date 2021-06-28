package hashmap

import "leetcode/utils"

func getIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	maps := make(map[*utils.ListNode]struct{}, 0)
	for ; headA != nil; headA = headA.Next {
		maps[headA] = struct{}{}
	}
	for ; headB != nil; headB = headB.Next {
		if _, ok := maps[headB]; ok {
			return headB
		}
	}
	return nil
}
