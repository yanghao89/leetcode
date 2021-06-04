package leetcode_543

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	fmt.Println(DiameterOfBinaryTree(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))
}
