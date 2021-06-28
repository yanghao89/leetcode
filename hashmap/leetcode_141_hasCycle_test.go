package hashmap

import "leetcode/utils"

func hasCycle(head *utils.ListNode) bool {
	maps := make(map[*utils.ListNode]struct{})
	for ; head != nil; head = head.Next {
		if _, ok := maps[head]; ok {
			return ok
		}
		maps[head] = struct{}{}
	}
	return false
}
