package leetcode_637

import (
	"leetcode/utils"
)

type data struct {
	Num   int
	Count int
}

func AverageOfLevels(root *utils.TreeNode) []float64 {
	levelData := []data{}
	var dfs func(root *utils.TreeNode, level int)
	dfs = func(node *utils.TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(levelData) {
			levelData[level].Num += root.Val
			levelData[level].Count++
		} else {
			levelData = append(levelData, data{node.Val, 1})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	arr := make([]float64, len(levelData))
	for i, d := range levelData {
		arr[i] = float64(d.Num) / float64(d.Count)
	}
	return arr
}
