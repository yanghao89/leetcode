package leetcode_653

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestFindTarget(t *testing.T) {
	node := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(FindTarget(node, 9))
}
