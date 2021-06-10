package leetcode_offer_236

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	fmt.Println(LowestCommonAncestor(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), utils.CreateTree([]int{3}), utils.CreateTree([]int{1})))
}
