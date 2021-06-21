package tree

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

var (
	tilt int
)

func TestFindTilt(t *testing.T) {
	fmt.Println(FindTilt(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))
}

func FindTilt(root *utils.TreeNode) int {
	tilt = 0
	helper(root)
	return tilt
}

func helper(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	//左节点的总坡度之和
	a := helper(root.Left)
	//右结点的总坡度之和
	b := helper(root.Right)
	//获取平均值
	tilt += int(math.Abs(float64(a - b)))
	//返回总坡度的总和
	return root.Val + a + b
}
