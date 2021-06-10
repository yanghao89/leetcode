package leetcode_100

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	node1 := utils.CreateTree([]int{1, 2, 3})
	node2 := node1
	fmt.Println(IsSameTree(node1, node2))
}
