package leetcode236

import "github.com/lucky51/algo/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return root
	}
	if rightNode == nil {
		return leftNode
	}
	return leftNode
}
