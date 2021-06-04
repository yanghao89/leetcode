package leetcode_617

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	fmt.Println(MergeTrees(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), utils.CreateTree([]int{10, 11, 12, 13, 14, 15, 16})))
}
