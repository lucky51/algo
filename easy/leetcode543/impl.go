package leetcode543

import "github.com/lucky51/algo/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *common.TreeNode) (ans int) {
	ans = 1
	var depth func(*common.TreeNode) int
	depth = func(tn *common.TreeNode) int {
		if tn == nil {
			return 0
		}
		l, r := depth(tn.Left), depth(tn.Right)
		ans = max(ans, l+r+1)
		return max(l, r) + 1
	}
	depth(root)
	ans -= 1
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
