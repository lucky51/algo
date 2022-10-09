package leetcode129

import . "github.com/lucky51/algo/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// sumNumbers 求根节点到叶节点数字之和
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, prevSum int) int {
	if node == nil {
		return 0
	}
	prevSum = prevSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return prevSum
	}
	return dfs(node.Left, prevSum) + dfs(node.Right, prevSum)
}
