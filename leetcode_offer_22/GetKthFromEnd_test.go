package leetcode_offer_22

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestGetKthFromEnd(t *testing.T) {
	node := utils.CreateListNode([]int{1, 1, 2, 3, 3})
	fmt.Println(GetKthFromEnd1(node, 2), GetKthFromEnd1(node, 3))
}
