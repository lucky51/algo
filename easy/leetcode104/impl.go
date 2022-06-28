package leetcode104

import "github.com/lucky51/algo/common"

func maxDepth(root *common.TreeNode) (ans int) {
	if root == nil {
		return
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
