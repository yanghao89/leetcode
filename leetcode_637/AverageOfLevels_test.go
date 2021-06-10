package leetcode_637

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	fmt.Println()
	node := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(AverageOfLevels(node))
}
