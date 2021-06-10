package leetcode_offer_22

import "leetcode/utils"

func GetKthFromEnd1(node *utils.ListNode, k int) *utils.ListNode {
	var res []*utils.ListNode
	for node != nil {
		res = append(res, node)
		node = node.Next
	}
	num := len(res)
	if num >= k {
		return res[num-k]
	}
	return nil
}

func GetKthFromEnd2(node *utils.ListNode, k int) *utils.ListNode {
	res := node
	for i := 1; node != nil; i++ {
		if i > k {
			res = res.Next
		}
		node = node.Next
	}
	return res
}
