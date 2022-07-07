package leetcode98

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
func isValidBST(root *common.TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *common.TreeNode, lower, super int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= super {
		return false
	}
	return helper(node.Left, lower, node.Val) && helper(node.Right, node.Val, super)
}
