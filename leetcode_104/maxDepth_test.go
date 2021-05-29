package leetcode_104

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	node1 := utils.BuildTree([]int{1, 2, 3, 4, 5})
	fmt.Println(MaxDepth(node1))
}
