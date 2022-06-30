package leetcode110

import (
	"github.com/lucky51/algo/common"
)

func isBalanced(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *common.TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
