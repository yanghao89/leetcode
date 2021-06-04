package leetcode_572

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestIsSubtree(t *testing.T) {
	fmt.Println(IsSubtree(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), utils.CreateTree([]int{1, 2, 3})))
}
