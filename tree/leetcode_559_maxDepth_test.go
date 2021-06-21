package tree

import (
	"leetcode/utils"
	"math"
	"testing"
)

func TestMaxDepth(t *testing.T) {

}
func MaxDepth(root *utils.Node) int {
	if root == nil {
		return 0
	}
	var l int
	for _, v := range root.Children {
		l = int(math.Max(float64(l), float64(MaxDepth(v))))
	}
	return l + 1
}
