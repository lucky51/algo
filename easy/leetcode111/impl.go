package leetcode111

import (
	"math"

	"github.com/lucky51/algo/common"
)

// minDepth 二叉树的最小深度
func minDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	mid := math.MaxInt32
	if root.Left != nil {
		mid = min(minDepth(root.Left), mid)
	}
	if root.Right != nil {
		mid = min(minDepth(root.Right), mid)
	}
	// 除了根和叶子结点，中间的部分 +1
	return mid + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
