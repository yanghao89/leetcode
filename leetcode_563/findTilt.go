package leetcode_563

import (
	"leetcode/utils"
	"math"
)

var (
	tilt int
)

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
