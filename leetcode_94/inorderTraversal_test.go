package leetcode_94

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	node := utils.BuildTree([]int{1, 2, 3})
	fmt.Println(InorderTraversal(node))
}
