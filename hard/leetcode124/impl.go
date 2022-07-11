package leetcode124

import (
	"math"

	"github.com/lucky51/algo/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// maxPathSum 二叉树中的最大路径和
func maxPathSum(root *common.TreeNode) int {
	maxPath := math.MinInt32
	var maxGain func(*common.TreeNode) int
	maxGain = func(tn *common.TreeNode) int {
		if tn == nil {
			return 0
		}
		// 只有大于0的子树才会被采用
		leftPath := max(maxGain(tn.Left), 0)
		rightPath := max(maxGain(tn.Right), 0)
		newMaxPath := leftPath + rightPath + tn.Val
		maxPath = max(newMaxPath, maxPath)
		// 单个节点计算最大值
		return tn.Val + max(leftPath, rightPath)
	}
	maxGain(root)
	return maxPath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
