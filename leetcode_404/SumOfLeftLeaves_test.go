package leetcode_404

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(SumOfLeftLeaves(root))
	fmt.Println(SumOfLeftLeaves01(root))
}
